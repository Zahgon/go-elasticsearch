package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _customNormalizer struct {
	v *types.CustomNormalizer
}

func NewCustomNormalizer() *_customNormalizer { _ = "STUB: not implemented"; return nil }

func (s *_customNormalizer) CharFilter(charfilters ...string) *_customNormalizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_customNormalizer) Filter(filters ...string) *_customNormalizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_customNormalizer) CustomNormalizerCaster() *types.CustomNormalizer {
	_ = "STUB: not implemented"
	return nil
}
