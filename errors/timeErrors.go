package errors

import "fmt"

type SeasonOutOfBoundsError struct {
	month int
}

func NewSeasonOutOfBoundsError(month int) SeasonOutOfBoundsError {
	return SeasonOutOfBoundsError{
		month: month,
	}
}

func (b SeasonOutOfBoundsError) Error() string {
	return fmt.Sprintf("Can't figure out season from month: %d. Valid values are in the range 1-8 inclusive.", b.month)
}
