package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _letterTokenizer struct {
	v *types.LetterTokenizer
}

func NewLetterTokenizer() *_letterTokenizer { _ = "STUB: not implemented"; return nil }

func (s *_letterTokenizer) Version(versionstring string) *_letterTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_letterTokenizer) LetterTokenizerCaster() *types.LetterTokenizer {
	_ = "STUB: not implemented"
	return nil
}
