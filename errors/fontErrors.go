package errors

import "fmt"

type Axis int

const (
	AxisHorizontal Axis = iota
	AxisVertical
)

func (a Axis) Line() string {
	switch a {
	case AxisHorizontal:
		return "row"
	case AxisVertical:
		return "column"
	default:
		return "<Error generating Error Message>"
	}
}

func (a Axis) Size() string {
	switch a {
	case AxisHorizontal:
		return "width"
	case AxisVertical:
		return "height"
	default:
		return "<Error generating Error Message>"
	}
}

type CharacterSizeError struct {
	axis      Axis
	wrongSize int
}

func NewCharacterSizeError(axis Axis, wrongSize int) CharacterSizeError {
	return CharacterSizeError{
		axis:      axis,
		wrongSize: wrongSize,
	}
}

func (c CharacterSizeError) Error() string {
	return fmt.Sprintf("Font needs a character %s of 1 or more. Not %d. Check the JSON files in assets/fonts/",
		c.axis.Size(), c.wrongSize)
}

type BigYShiftError struct {
	yShift     int
	charHeight int
}

func NewBigYShiftError(yshift int, height int) BigYShiftError {
	return BigYShiftError{
		yShift:     yshift,
		charHeight: height,
	}
}

func (b BigYShiftError) Error() string {
	return fmt.Sprintf("Font's Y-Shift(%d) places all the characters lower than the basline as it's larger than the character height(%d). Check the JSON files in assets/fonts/",
		b.yShift, b.charHeight)
}

type IncorrectLineCountError struct {
	axis         Axis
	lineCount    int
	correctCount int
}

func NewIncorrectLineCountError(axis Axis, wrong int, correct int) IncorrectLineCountError {
	return IncorrectLineCountError{
		axis:         axis,
		lineCount:    wrong,
		correctCount: correct,
	}
}

func (i IncorrectLineCountError) Error() string {
	return fmt.Sprintf("Font has the wrong number of %ss in 'mapping' than defined by '%ss'. Found %d, Expected %d. Check the JSON files in assets/fonts/",
		i.axis.Line(), i.axis.Line(), i.lineCount, i.correctCount)
}

type RepeatedCharacterDefinitionError struct {
	character string
}

func NewRepeatedCharacterDefinitionError(char string) RepeatedCharacterDefinitionError {
	return RepeatedCharacterDefinitionError{
		character: char,
	}
}

func (r RepeatedCharacterDefinitionError) Error() string {
	return fmt.Sprintf("Font can't have multiple of %v in the definition. Check the JSON files in assets/fonts/",
		r.character)
}

type InvalidCharacterError struct {
	character string
}

func NewInvalidCharacterError(char string) InvalidCharacterError {
	return InvalidCharacterError{
		character: char,
	}
}

func (i InvalidCharacterError) Error() string {
	return fmt.Sprintf("%s isn't a valid unicode character or 'unknown'. Check the JSON files in assets/fonts/",
		i.character)
}

type MissingUnknownGlyphError struct {
}

func NewMissingUnknownError() MissingUnknownGlyphError {
	return MissingUnknownGlyphError{}
}

func (m MissingUnknownGlyphError) Error() string {
	return "Font requires an 'unknown' glyph to represent missing characters. Add one to the JSON files in assets/fonts/"
}

type RuneFromUnloadedResourcesError struct {
}

func NewRuneFromUnloadedResourcesError() RuneFromUnloadedResourcesError {
	return RuneFromUnloadedResourcesError{}
}

func (m RuneFromUnloadedResourcesError) Error() string {
	return "Can't get a rune from a font before calling `UpdateResources` on it atleast once."
}
