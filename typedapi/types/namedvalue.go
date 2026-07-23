package types

type NamedValue map[string][]FieldValue

type NamedValueVariant interface {
	NamedValueCaster() *NamedValue
}
