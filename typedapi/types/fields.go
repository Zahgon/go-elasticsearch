package types

type Fields []string

type FieldsVariant interface {
	FieldsCaster() *Fields
}
