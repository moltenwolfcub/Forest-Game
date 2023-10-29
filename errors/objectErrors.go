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

type UnknownPumpkinVariantError struct {
	variant string
}

func NewUnknownPumpkinVariantError(variant string) UnknownPumpkinVariantError {
	return UnknownPumpkinVariantError{
		variant: variant,
	}
}

func (p UnknownPumpkinVariantError) Error() string {
	return fmt.Sprintf("Unkown pumpkin variant: %s", p.variant)
}

type InvalidPumpkinPhaseError struct {
	phase string
}

func NewInvalidPumpkinPhaseError(phase string) InvalidPumpkinPhaseError {
	return InvalidPumpkinPhaseError{
		phase: phase,
	}
}

func (b InvalidPumpkinPhaseError) Error() string {
	return fmt.Sprintf("Not a valid pumpkin phase: %s", b.phase)
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
