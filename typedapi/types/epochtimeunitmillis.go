package types

type EpochTimeUnitMillis int64

type EpochTimeUnitMillisVariant interface {
	EpochTimeUnitMillisCaster() *EpochTimeUnitMillis
}
