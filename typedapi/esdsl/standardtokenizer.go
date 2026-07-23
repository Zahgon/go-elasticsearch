package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _standardTokenizer struct {
	v *types.StandardTokenizer
}

func NewStandardTokenizer() *_standardTokenizer { _ = "STUB: not implemented"; return nil }

func (s *_standardTokenizer) MaxTokenLength(maxtokenlength int) *_standardTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_standardTokenizer) Version(versionstring string) *_standardTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_standardTokenizer) StandardTokenizerCaster() *types.StandardTokenizer {
	_ = "STUB: not implemented"
	return nil
}
