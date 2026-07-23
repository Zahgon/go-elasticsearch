package types

type ScalarValue any

type ScalarValueVariant interface {
	ScalarValueCaster() *ScalarValue
}
