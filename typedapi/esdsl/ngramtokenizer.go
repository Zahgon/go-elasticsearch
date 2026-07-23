package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/tokenchar"
)

type _nGramTokenizer struct {
	v *types.NGramTokenizer
}

func NewNGramTokenizer() *_nGramTokenizer { _ = "STUB: not implemented"; return nil }

func (s *_nGramTokenizer) CustomTokenChars(customtokenchars string) *_nGramTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nGramTokenizer) MaxGram(maxgram int) *_nGramTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nGramTokenizer) MinGram(mingram int) *_nGramTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nGramTokenizer) TokenChars(tokenchars ...tokenchar.TokenChar) *_nGramTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nGramTokenizer) Version(versionstring string) *_nGramTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nGramTokenizer) NGramTokenizerCaster() *types.NGramTokenizer {
	_ = "STUB: not implemented"
	return nil
}
