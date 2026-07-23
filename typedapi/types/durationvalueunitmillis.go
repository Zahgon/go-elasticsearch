package types

type DurationValueUnitMillis int64

type DurationValueUnitMillisVariant interface {
	DurationValueUnitMillisCaster() *DurationValueUnitMillis
}
