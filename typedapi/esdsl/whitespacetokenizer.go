package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _whitespaceTokenizer struct {
	v *types.WhitespaceTokenizer
}

func NewWhitespaceTokenizer() *_whitespaceTokenizer { _ = "STUB: not implemented"; return nil }

func (s *_whitespaceTokenizer) MaxTokenLength(maxtokenlength int) *_whitespaceTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_whitespaceTokenizer) Version(versionstring string) *_whitespaceTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_whitespaceTokenizer) WhitespaceTokenizerCaster() *types.WhitespaceTokenizer {
	_ = "STUB: not implemented"
	return nil
}
