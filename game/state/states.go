package state

import "fmt"

type State struct {
	properties []Property
}

func (s State) GetProperty(str string) *Property {
	for i, p := range s.properties {
		if p.matchesName(str) {
			return &s.properties[i]
		}
	}
	panic(fmt.Sprintf("Couldn't get property %s from %s", str, s))
}
func (s State) GetValue(str string) any {
	prop := s.GetProperty(str)
	return prop.getValue()
}

func (s State) UpdateValue(str string, val any) {
	prop := s.GetProperty(str)
	prop.setValue(val)
}

func (s State) String() string {
	str := "State["
	length := len(s.properties)
	i := 1

	for _, k := range s.properties {
		str += k.String()
		if i < length {
			str += ", "
		}
		i++
	}
	return str + "]"
}

type StateBuilder struct {
	properties []Property
}

func StateBuilderFromStr(str string) StateBuilder {
	panic("Not Implemented")
}

func (s *StateBuilder) Add(prop ...Property) *StateBuilder {
	s.properties = append(s.properties, prop...)
	return s
}

func (s StateBuilder) Build() State {
	newState := State(s)
	return newState
}

// type BerryState struct {
// 	age     IntegerProperty
// 	variant EnumProperty[BerryVariant]
// }

// type ObjectState struct {
// }

// type StateDefinition[O any] struct {
// 	owner            O
// 	propertiesByName map[string]any
// 	states           []ObjectState
// 	prop             Property[any]
// }

// type StateDefinitionBuilder struct {
// }

// type SyntheticBerry struct {
// 	stateDefinition StateDefinition[Berry]

// 	age IntegerProperty
// 	// variant EnumProperty[BerryVariant]
// }

// func (obj SyntheticBerry) RegisterStateDef(*StateDefinitionBuilder) {
// 	S := StateDefinition[any]{
// 		prop: NewIntegerProperty("test", 0, 5),
// 	}
// }
