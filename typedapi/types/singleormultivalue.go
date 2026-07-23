package types

type SingleOrMultiValue []FieldValue

type SingleOrMultiValueVariant interface {
	SingleOrMultiValueCaster() *SingleOrMultiValue
}
