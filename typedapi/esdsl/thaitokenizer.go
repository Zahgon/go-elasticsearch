package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _thaiTokenizer struct {
	v *types.ThaiTokenizer
}

func NewThaiTokenizer() *_thaiTokenizer { _ = "STUB: not implemented"; return nil }

func (s *_thaiTokenizer) Version(versionstring string) *_thaiTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_thaiTokenizer) ThaiTokenizerCaster() *types.ThaiTokenizer {
	_ = "STUB: not implemented"
	return nil
}
