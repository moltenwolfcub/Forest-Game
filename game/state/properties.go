package state

import (
	"fmt"
	"strings"
)

type Property struct {
	value string
	name  string
}

func NewProperty(name string, Val string) Property {
	return Property{
		name:  name,
		value: Val,
	}
}
func PropertyFromString(str string) Property {
	parts := strings.Split(str, "=")
	partLen := len(parts)
	if partLen < 2 {
		panic(fmt.Sprintf("Incorrect property string: %s. Properties should be in the format of <Name>=<Value>", str))
	}
	if partLen > 2 {
		panic(fmt.Sprintf("Only one equals sign should be used in a property string: %s. Properties should be in the format of <Name>=<Value>", str))
	}

	return NewProperty(strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]))
}

func (s Property) String() string {
	return fmt.Sprintf("%s=%v", s.name, s.value)
}

func (p Property) matchesName(str string) bool {
	return p.name == str
}

func (p Property) getValue() string {
	return p.value
}

func (p *Property) setValue(val string) {
	p.value = val
}

// func (p Property) GetValueImproved[T type]() T {
// 	val, ok := p.value.(int)
// 	if !ok {
// 		panic(fmt.Sprintf("They value stored in %s isn't a %T", p, T))
// 	}
// 	return val
// }

// func (p Property) GetIntValue() int {
// 	if p.valueType.Kind() != reflect.Int {
// 		panic("Not an integer property")
// 	}

// 	return p.value.(int)
// }

//========== OLD ==========

// type Property[T any] interface {
// 	GetPossibleValues() []T
// 	GetName() string
// 	ConvertToString(T) string
// 	ConvertToValue(string) T
// }

// // implements Property[int]
// type IntegerProperty struct {
// 	Name   string
// 	min    int
// 	max    int
// 	values []int
// }

// func (p IntegerProperty) GetPossibleValues() []int {
// 	return p.values
// }

// func (p IntegerProperty) GetName() string {
// 	return p.Name
// }

// func (p IntegerProperty) ConvertToString(value int) string {
// 	return fmt.Sprintf("%d", value)
// }

// func (p IntegerProperty) ConvertToValue(value string) int {
// 	i, err := strconv.Atoi(value)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return i
// }

// func NewIntegerProperty(name string, min int, max int) IntegerProperty {
// 	if min < 0 {
// 		panic(fmt.Sprintf("Min value of %s must be 0 or bigger", name))
// 	}
// 	if max <= min {
// 		panic(fmt.Sprintf("Max value of %s must be bigger than the min (%d)", name, min))
// 	}
// 	vals := []int{}
// 	for i := min; i <= max; i++ {
// 		vals = append(vals, i)
// 	}
// 	property := IntegerProperty{
// 		min:    min,
// 		max:    max,
// 		values: vals,
// 	}
// 	return property
// }

// // implements Property[bool]
// type BooleanProperty struct {
// 	Name string
// 	// Value bool
// }

// // implements Property[~int]
// type EnumProperty[Enum ~int] struct {
// 	Name string
// 	// Value Enum
// }
