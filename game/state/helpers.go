package state

import (
	"fmt"
	"strconv"
)

func GetIntFromState[T ~int](s State, valName string) T {
	valueStr := s.GetValue(valName)
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		panic(fmt.Sprintf("Cannot convert '%s' property to an integer", valName))
	}

	finalValue := T(value)
	return finalValue
}
