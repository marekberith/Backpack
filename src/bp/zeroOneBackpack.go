package bp

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Pair struct {
	price, weight int
}

type TwoDimMatrixValue struct {
	Value    int
	from     *Pair
	included bool
}

func parseMatrix(tuples string, weight, itemsCount int) ([][]TwoDimMatrixValue, []Pair, error) {
	tuples = strings.Replace(tuples, " ", "", -1)
	splittedTuples := strings.Split(tuples, ";")
	if weight < 0 {
		return nil, nil, errors.New("invalid size of the backpack")
	}
	if len(splittedTuples) != itemsCount {
		return nil, nil, errors.New("too few/many tuples entered")
	}
	pairs := make([]Pair, len(splittedTuples))
	matrix := make([][]TwoDimMatrixValue, len(splittedTuples)+1)
	for i, raw := range splittedTuples {
		matrix[i] = make([]TwoDimMatrixValue, weight+1)
		if !(raw[0] == '(' && raw[len(raw)-1] == ')') {
			return nil, nil, errors.New("invalid data")
		}
		n, err := fmt.Sscanf(raw[1:len(raw)-1], "%d,%d", &pairs[i].price, &pairs[i].weight)
		if n != 2 && err != nil {
			return nil, nil, errors.New("invalid data")
		}
	}
	matrix[len(splittedTuples)] = make([]TwoDimMatrixValue, weight+1)
	return matrix, pairs, nil
}

func max(fVal, sVal int, fPair, sPair *Pair) TwoDimMatrixValue {
	if fVal > sVal {
		return TwoDimMatrixValue{fVal, fPair, false}
	}
	return TwoDimMatrixValue{sVal, sPair, true}
}

type TwoDimMatrix [][]TwoDimMatrixValue

func (m TwoDimMatrix) get(i, j, price int) int {
	if i < 0 || j < 0 {
		return -1
	}
	return m[i][j].Value + price
}

func calculateMatrix(matrix [][]TwoDimMatrixValue, pairs []Pair, weight int) ([][]TwoDimMatrixValue, *TwoDimMatrixValue, *Pair) {
	entriesCount := len(matrix)
	var maxValue *TwoDimMatrixValue
	var maxPair *Pair
	for i := 1; i < entriesCount; i++ {
		for j := 1; j < weight+1; j++ {
			weightWithout := j - pairs[i-1].weight
			matrix[i][j] = max(TwoDimMatrix(matrix).get(i-1, j, 0), TwoDimMatrix(matrix).get(i-1, weightWithout, pairs[i-1].price), &Pair{i - 1, j}, &Pair{i - 1, weightWithout})
			if maxValue == nil || maxValue.Value < matrix[i][j].Value {
				maxValue = &matrix[i][j]
				maxPair = &Pair{i, j}
			}
		}
	}
	return matrix, maxValue, maxPair
}

func getZOBackpackResult(matrix TwoDimMatrix, previous *TwoDimMatrixValue, previousPair *Pair) (TwoDimMatrix, string) {
	if previousPair == nil {
		return nil, "No solution exists"
	}
	result := "One of the solutions: ("
	for previous != nil && previousPair != nil {
		if previous.included == true {
			result += "(" + strconv.Itoa(previousPair.price) + ", " + strconv.Itoa(previousPair.weight) + ")"
		}
		previousPair = previous.from
		if previousPair == nil || previousPair.price == 0 || previousPair.weight == 0 {
			break
		}
		if previous.included == true {
			result += ", "
		}
		previous = &matrix[previousPair.price][previousPair.weight]
	}
	result += ")"
	return matrix, result
}

func SolveZeroOneBackpack(tuples string, weight, itemsCount int) (TwoDimMatrix, string) {
	emptyMatrix, pairs, err := parseMatrix(tuples, weight, itemsCount)
	if err != nil {
		return nil, fmt.Sprintf("%s", err)
	}
	matrix, maxValue, maxPair := calculateMatrix(emptyMatrix, pairs, weight)
	return getZOBackpackResult(matrix, maxValue, maxPair)
}
