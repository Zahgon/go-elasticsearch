package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _normalizer struct {
	v types.Normalizer
}

func NewNormalizer() *_normalizer { _ = "STUB: not implemented"; return nil }

func (u *_normalizer) LowercaseNormalizer(lowercasenormalizer types.LowercaseNormalizerVariant) *_normalizer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_lowercaseNormalizer) NormalizerCaster() *types.Normalizer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_normalizer) CustomNormalizer(customnormalizer types.CustomNormalizerVariant) *_normalizer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_customNormalizer) NormalizerCaster() *types.Normalizer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_normalizer) NormalizerCaster() *types.Normalizer { _ = "STUB: not implemented"; return nil }
