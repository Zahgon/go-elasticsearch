package types

type Validation any

type ValidationVariant interface {
	ValidationCaster() *Validation
}
