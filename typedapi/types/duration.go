package types

type Duration any

type DurationVariant interface {
	DurationCaster() *Duration
}
