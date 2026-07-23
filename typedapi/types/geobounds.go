package types

type GeoBounds any

type GeoBoundsVariant interface {
	GeoBoundsCaster() *GeoBounds
}
