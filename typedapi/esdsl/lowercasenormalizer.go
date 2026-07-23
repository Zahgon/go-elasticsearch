package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _lowercaseNormalizer struct {
	v *types.LowercaseNormalizer
}

func NewLowercaseNormalizer() *_lowercaseNormalizer { _ = "STUB: not implemented"; return nil }

func (s *_lowercaseNormalizer) LowercaseNormalizerCaster() *types.LowercaseNormalizer {
	_ = "STUB: not implemented"
	return nil
}
