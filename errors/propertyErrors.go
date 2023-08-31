package errors

import "fmt"

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
