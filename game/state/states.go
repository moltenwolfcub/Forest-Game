package state

import (
	"fmt"
	"sort"
	"strings"
)

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
func (s State) GetValue(str string) string {
	prop := s.GetProperty(str)
	return prop.getValue()
}

func (s State) UpdateValue(str string, val string) {
	prop := s.GetProperty(str)
	prop.setValue(val)
}

func (s State) ToTextureKey() string {
	sorted := s.properties
	sort.SliceStable(sorted, func(i, j int) bool {
		names := []string{
			sorted[i].name,
			sorted[j].name,
		}

		sort.Strings(names)
		return names[0] == sorted[i].name
	})

	str := ""
	length := len(sorted)
	i := 1

	for _, k := range sorted {
		str += k.String()
		if i < length {
			str += ", "
		}
		i++
	}
	return str
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
	builder := StateBuilder{}

	props := strings.Split(str, ",")
	for _, str := range props {
		builder.Add(PropertyFromString(strings.TrimSpace(str)))
	}
	return builder
}

func (s *StateBuilder) Add(prop ...Property) *StateBuilder {
	s.properties = append(s.properties, prop...)
	return s
}

func (s StateBuilder) Build() State {
	newState := State(s)
	return newState
}
