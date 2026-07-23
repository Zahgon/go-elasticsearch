package types

type Missing any

type MissingVariant interface {
	MissingCaster() *Missing
}
