package errors

import "fmt"

type UnknownBerryVariantError struct {
	variant string
}

func NewUnknownBerryVariantError(variant string) UnknownBerryVariantError {
	return UnknownBerryVariantError{
		variant: variant,
	}
}

func (b UnknownBerryVariantError) Error() string {
	return fmt.Sprintf("Unkown berry variant: %s", b.variant)
}

type InvalidBerryPhaseError struct {
	phase string
}

func NewInvalidBerryPhaseError(phase string) InvalidBerryPhaseError {
	return InvalidBerryPhaseError{
		phase: phase,
	}
}

func (b InvalidBerryPhaseError) Error() string {
	return fmt.Sprintf("Not a valid berry phase: %s", b.phase)
}

type MultiHitboxRiverSegmentError struct {
	hitboxCount int
}

func NewMultiHitboxRiverSegmentError(hitboxCount int) MultiHitboxRiverSegmentError {
	return MultiHitboxRiverSegmentError{
		hitboxCount: hitboxCount,
	}
}

func (m MultiHitboxRiverSegmentError) Error() string {
	return fmt.Sprintf("A river segment should only have a single rect for a hitbox not %d. If more are required use a river with multiple segments.", m.hitboxCount)
}

type UnknownMushroomVariantError struct {
	variant string
}

func NewUnknownMushroomVariantError(variant string) UnknownMushroomVariantError {
	return UnknownMushroomVariantError{
		variant: variant,
	}
}

func (m UnknownMushroomVariantError) Error() string {
	return fmt.Sprintf("Unkown mushroom variant: %s", m.variant)
}

type InvalidMushroomPhaseError struct {
	phase string
}

func NewInvalidMushroomPhaseError(phase string) InvalidMushroomPhaseError {
	return InvalidMushroomPhaseError{
		phase: phase,
	}
}

func (b InvalidMushroomPhaseError) Error() string {
	return fmt.Sprintf("Not a valid mushroom phase: %s", b.phase)
}
