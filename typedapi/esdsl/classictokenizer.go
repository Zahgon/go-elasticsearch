package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _classicTokenizer struct {
	v *types.ClassicTokenizer
}

func NewClassicTokenizer() *_classicTokenizer { _ = "STUB: not implemented"; return nil }

func (s *_classicTokenizer) MaxTokenLength(maxtokenlength int) *_classicTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_classicTokenizer) Version(versionstring string) *_classicTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_classicTokenizer) ClassicTokenizerCaster() *types.ClassicTokenizer {
	_ = "STUB: not implemented"
	return nil
}
