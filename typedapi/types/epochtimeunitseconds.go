package types

type EpochTimeUnitSeconds int64

type EpochTimeUnitSecondsVariant interface {
	EpochTimeUnitSecondsCaster() *EpochTimeUnitSeconds
}
