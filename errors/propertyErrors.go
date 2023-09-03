package errors

import (
	"fmt"
	"reflect"
)

type UnknownTextureKeyError struct {
	textureKey  string
	knownStates []string
}

func NewUnknownTextureKeyError(key string, known []string) UnknownTextureKeyError {
	return UnknownTextureKeyError{
		textureKey:  key,
		knownStates: known,
	}
}

func (u UnknownTextureKeyError) Error() string {
	return fmt.Sprintf("Can't handle unknown texture key: %s\n\nKnown states are: %v", u.textureKey, u.knownStates)
}

type UnobtainablePropertyError struct {
	state    string
	property string
}

func NewUnobtainablePropertyError(state string, property string) UnobtainablePropertyError {
	return UnobtainablePropertyError{
		state:    state,
		property: property,
	}
}

func (u UnobtainablePropertyError) Error() string {
	return fmt.Sprintf("Couldn't get property %s from %s", u.property, u.state)
}

type PropertyConversionError struct {
	propType reflect.Type
	propName string
}

func NewPropertyConversionError(propertyType reflect.Type, value string) PropertyConversionError {
	return PropertyConversionError{
		propType: propertyType,
		propName: value,
	}
}

func (p PropertyConversionError) Error() string {
	return fmt.Sprintf("Cannot convert '%s' property to an %v", p.propName, p.propType)
}

type BadPropertyStringError struct {
	bad string
}

func NewBadPropertyStringError(badStr string) BadPropertyStringError {
	return BadPropertyStringError{
		bad: badStr,
	}
}

func (b BadPropertyStringError) Error() string {
	return fmt.Sprintf("Incorrect property string: %s. Properties should be in the format of <Name>=<Value>", b.bad)
}
