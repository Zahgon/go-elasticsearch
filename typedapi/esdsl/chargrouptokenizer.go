package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _charGroupTokenizer struct {
	v *types.CharGroupTokenizer
}

func NewCharGroupTokenizer() *_charGroupTokenizer { _ = "STUB: not implemented"; return nil }

func (s *_charGroupTokenizer) MaxTokenLength(maxtokenlength int) *_charGroupTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_charGroupTokenizer) TokenizeOnChars(tokenizeonchars ...string) *_charGroupTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_charGroupTokenizer) Version(versionstring string) *_charGroupTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_charGroupTokenizer) CharGroupTokenizerCaster() *types.CharGroupTokenizer {
	_ = "STUB: not implemented"
	return nil
}
