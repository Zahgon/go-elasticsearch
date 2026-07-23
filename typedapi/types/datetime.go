package types

type DateTime any

type DateTimeVariant interface {
	DateTimeCaster() *DateTime
}
