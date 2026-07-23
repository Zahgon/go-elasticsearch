package types

type FieldDateMath any

type FieldDateMathVariant interface {
	FieldDateMathCaster() *FieldDateMath
}
