package types

type CharFilter any

type CharFilterVariant interface {
	CharFilterCaster() *CharFilter
}
