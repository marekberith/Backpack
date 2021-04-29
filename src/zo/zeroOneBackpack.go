package zo

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Pair struct {
	price, weight int
}

type MatrixValue struct {
	Value int
	from  *Pair
}

func parseMatrix(tuples string, weight, itemsCount int) ([][]MatrixValue, []Pair, error) {
	tuples = strings.Replace(tuples, " ", "", -1)
	splittedTuples := strings.Split(tuples, ";")
	if len(splittedTuples) != itemsCount {
		return nil, nil, errors.New("too few/many tuples entered")
	}
	pairs := make([]Pair, len(splittedTuples))
	matrix := make([][]MatrixValue, len(splittedTuples) + 1)
	for i, raw := range splittedTuples {
		matrix[i] = make([]MatrixValue, weight + 1)
		if !(raw[0] == '(' && raw[len(raw) - 1] == ')') {
			return nil, nil, errors.New("invalid data")
		}
		n, err := fmt.Sscanf(raw[1:len(raw)-1], "%d,%d", &pairs[i].price, &pairs[i].weight)
		if n != 2 && err != nil {
			return nil, nil, errors.New("invalid data")
		}
	}
	matrix[len(splittedTuples)] = make([]MatrixValue, weight + 1)
	return matrix, pairs, nil
}

func max(fVal, sVal int, fPair, sPair *Pair) (int, *Pair) {
	if fVal > sVal {
		return fVal, fPair
	}
	return sVal, sPair
}

type Matrix [][]MatrixValue

func (m Matrix) get (i, j, price int) int{
	if i < 0 || j < 0 {
		return -1
	}
	return m[i][j].Value + price
}

func calculateMatrix(matrix [][]MatrixValue, pairs []Pair, weight int) ([][]MatrixValue, string){
	entriesCount := len(matrix)
	var maxValue *MatrixValue
	var maxPair *Pair
	for i := 1; i < entriesCount; i++ {
		for j := 1; j < weight + 1; j++ {
			weightWithout := j - pairs[i - 1].weight
			matrix[i][j].Value, matrix[i][j].from = max(Matrix(matrix).get(i - 1, j, 0), Matrix(matrix).get(i - 1, weightWithout, pairs[i - 1].price), &Pair{i - 1,j}, &Pair{i - 1, weightWithout})
			if maxValue == nil || maxValue.Value < matrix[i][j].Value {
				maxValue = &matrix[i][j]
				maxPair = &Pair{i, j}
			}
		}
	}
	previousPair := maxPair
	previous := maxValue
	result := "One of the solutions: ("
	for previous != nil && previousPair != nil {
		result += "(" + strconv.Itoa(previousPair.price) + ", " + strconv.Itoa(previousPair.weight) + ")"
		previousPair = previous.from
		if previousPair == nil || previousPair.price == 0 || previousPair.weight == 0 {
			result += ")"
			break
		}
		result += ", "
		previous = &matrix[previousPair.price][previousPair.weight]
	}
	return matrix, result
}

func GetZeroOneMatrix(tuples string, weight, itemsCount int) (Matrix, string) {
	emptyMatrix, pairs, err := parseMatrix(tuples, weight, itemsCount)
	if err != nil {
		return nil, fmt.Sprintf("%s", err)
	}
	matrix, result := calculateMatrix(emptyMatrix, pairs, weight)
	return matrix, result
}