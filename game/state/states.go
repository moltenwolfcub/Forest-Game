package state

import (
	"sort"
	"strings"

	"github.com/moltenwolfcub/Forest-Game/errors"
)

type State struct {
	properties []Property
}

func (s State) GetProperty(str string) (*Property, error) {
	for i, p := range s.properties {
		if p.matchesName(str) {
			return &s.properties[i], nil
		}
	}
	return nil, errors.NewUnobtainablePropertyError(s.String(), str)
}
func (s State) GetValue(str string) (string, error) {
	prop, err := s.GetProperty(str)
	if err != nil {
		return "", err
	}
	return prop.getValue(), nil
}

func (s State) UpdateValue(str string, val string) error {
	prop, err := s.GetProperty(str)
	if err != nil {
		return err
	}

	prop.setValue(val)
	return nil
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

func StateBuilderFromStr(str string) (StateBuilder, error) {
	builder := StateBuilder{}

	props := strings.Split(str, ",")
	for _, str := range props {
		property, err := PropertyFromString(strings.TrimSpace(str))
		if err != nil {
			return StateBuilder{}, err
		}

		builder.Add(property)
	}
	return builder, nil
}

func (s *StateBuilder) Add(prop ...Property) *StateBuilder {
	s.properties = append(s.properties, prop...)
	return s
}

func (s StateBuilder) Build() State {
	newState := State(s)
	return newState
}
