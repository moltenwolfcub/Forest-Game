package state

import (
	"reflect"
	"strconv"

	"github.com/moltenwolfcub/Forest-Game/errors"
)

func GetIntFromState[T ~int](s State, valName string) (T, error) {
	valueStr, err := s.GetValue(valName)
	if err != nil {
		return 0, err
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return 0, errors.NewPropertyConversionError(reflect.TypeOf(0), valName)
	}

	finalValue := T(value)
	return finalValue, nil
}
