package types

type Property any

type PropertyVariant interface {
	PropertyCaster() *Property
}
