package types

type FieldValue any

type FieldValueVariant interface {
	FieldValueCaster() *FieldValue
}
