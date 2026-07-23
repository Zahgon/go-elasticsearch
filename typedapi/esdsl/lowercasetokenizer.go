package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _lowercaseTokenizer struct {
	v *types.LowercaseTokenizer
}

func NewLowercaseTokenizer() *_lowercaseTokenizer { _ = "STUB: not implemented"; return nil }

func (s *_lowercaseTokenizer) Version(versionstring string) *_lowercaseTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_lowercaseTokenizer) LowercaseTokenizerCaster() *types.LowercaseTokenizer {
	_ = "STUB: not implemented"
	return nil
}
