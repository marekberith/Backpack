package bp

import (
	"errors"
	"fmt"
	"strings"
)

func parseVectorUnbounded(tuples string, weight, itemsCount int) ([]VectorValue, []Pair, error) {
	tuples = strings.Replace(tuples, " ", "", -1)
	splittedTuples := strings.Split(tuples, ";")
	if weight < 0 {
		return nil, nil, errors.New("invalid size of the backpack")
	}
	if len(splittedTuples) != itemsCount {
		return nil, nil, errors.New("too few/many tuples entered")
	}
	pairs := make([]Pair, len(splittedTuples))
	vector := make([]VectorValue, weight+1)
	for i, raw := range splittedTuples {
		if !(raw[0] == '(' && raw[len(raw)-1] == ')') {
			return nil, nil, errors.New("invalid data")
		}
		n, err := fmt.Sscanf(raw[1:len(raw)-1], "%d,%d", &pairs[i].price, &pairs[i].weight)
		if n != 2 && err != nil {
			return nil, nil, errors.New("invalid data")
		}
	}
	return vector, pairs, nil
}

func SolveUnboundedBackpack(tuples string, weight, itemsCount int) (Vector, string) {
	emptyVector, pairs, err := parseVectorUnbounded(tuples, weight, itemsCount)
	if err != nil {
		return nil, fmt.Sprintf("%s", err)
	}
	calculateVectorUnbounded(emptyVector, pairs)
	return getUnboundedBackpackResult(emptyVector)
}

func getUnboundedBackpackResult(vector Vector) (Vector, string) {
	if vector[len(vector)-1].Value == 0 {
		return nil, "No solution exists"
	}
	previous := &vector[len(vector)-1]
	usedItems := make(map[Pair]int)
	for previous != nil {
		value, exists := usedItems[*previous.using]
		if exists {
			usedItems[*previous.using] = value + 1
		} else {
			usedItems[*previous.using] = 1
		}
		if previous.from == 0 {
			break
		}
		previous = &vector[previous.from]
	}
	result := fmt.Sprintf("Price: %d, Number of items: %d, ", vector[len(vector)-1].Value, len(usedItems))
	result += "Items: ("
	for key, value := range usedItems {
		result += fmt.Sprintf("(%d, %d, %dx),", key.price, key.weight, value)
	}
	result = result[:len(result)-1] + ")"
	return vector, result
}

type VectorValue struct {
	Value int
	using *Pair
	from  int
}

type Vector []VectorValue

func (o Vector) get(i int) int {
	if i < 0 {
		return -1
	}
	return o[i].Value
}

func (o Vector) maxUnbounded(pairs []Pair, position int) VectorValue {
	var maxVal VectorValue
	for i, pair := range pairs {
		movePrice := o.get(position - pair.weight)
		if movePrice == -1 {
			continue
		}
		pairVal := pair.price + movePrice
		if pairVal > maxVal.Value {
			maxVal.Value = pairVal
			maxVal.using = &pairs[i]
			maxVal.from = position - pair.weight
		}
	}
	return maxVal
}

func calculateVectorUnbounded(vector []VectorValue, pairs []Pair) {
	length := len(vector)
	for i := 1; i < length; i++ {
		vector[i] = Vector(vector).maxUnbounded(pairs, i)
	}
}
