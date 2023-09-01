package state

import (
	"fmt"
	"strings"

	"github.com/moltenwolfcub/Forest-Game/errors"
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
func PropertyFromString(str string) (Property, error) {
	parts := strings.Split(str, "=")
	partLen := len(parts)
	if partLen != 2 {
		return Property{}, errors.NewBadPropertyStringError(str)
	}

	return NewProperty(strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])), nil
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
