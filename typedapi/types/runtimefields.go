package types

type RuntimeFields map[string]RuntimeField

type RuntimeFieldsVariant interface {
	RuntimeFieldsCaster() *RuntimeFields
}
