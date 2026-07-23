package types

type Normalizer any

type NormalizerVariant interface {
	NormalizerCaster() *Normalizer
}
