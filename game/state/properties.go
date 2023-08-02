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
