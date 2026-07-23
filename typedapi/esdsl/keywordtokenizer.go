package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _keywordTokenizer struct {
	v *types.KeywordTokenizer
}

func NewKeywordTokenizer() *_keywordTokenizer { _ = "STUB: not implemented"; return nil }

func (s *_keywordTokenizer) BufferSize(buffersize int) *_keywordTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_keywordTokenizer) Version(versionstring string) *_keywordTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_keywordTokenizer) KeywordTokenizerCaster() *types.KeywordTokenizer {
	_ = "STUB: not implemented"
	return nil
}
