package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/tokenchar"
)

type _edgeNGramTokenizer struct {
	v *types.EdgeNGramTokenizer
}

func NewEdgeNGramTokenizer() *_edgeNGramTokenizer { _ = "STUB: not implemented"; return nil }

func (s *_edgeNGramTokenizer) CustomTokenChars(customtokenchars string) *_edgeNGramTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_edgeNGramTokenizer) MaxGram(maxgram int) *_edgeNGramTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_edgeNGramTokenizer) MinGram(mingram int) *_edgeNGramTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_edgeNGramTokenizer) TokenChars(tokenchars ...tokenchar.TokenChar) *_edgeNGramTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_edgeNGramTokenizer) Version(versionstring string) *_edgeNGramTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_edgeNGramTokenizer) EdgeNGramTokenizerCaster() *types.EdgeNGramTokenizer {
	_ = "STUB: not implemented"
	return nil
}
