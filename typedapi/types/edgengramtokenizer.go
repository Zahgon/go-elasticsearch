package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/tokenchar"
)

type EdgeNGramTokenizer struct {
	CustomTokenChars *string               `json:"custom_token_chars,omitempty"`
	MaxGram          *int                  `json:"max_gram,omitempty"`
	MinGram          *int                  `json:"min_gram,omitempty"`
	TokenChars       []tokenchar.TokenChar `json:"token_chars,omitempty"`
	Type             string                `json:"type,omitempty"`
	Version          *string               `json:"version,omitempty"`
}

func (s *EdgeNGramTokenizer) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s EdgeNGramTokenizer) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewEdgeNGramTokenizer() *EdgeNGramTokenizer { _ = "STUB: not implemented"; return nil }

type EdgeNGramTokenizerVariant interface {
	EdgeNGramTokenizerCaster() *EdgeNGramTokenizer
}

func (s *EdgeNGramTokenizer) EdgeNGramTokenizerCaster() *EdgeNGramTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *EdgeNGramTokenizer) TokenizerDefinitionCaster() *TokenizerDefinition {
	_ = "STUB: not implemented"
	return nil
}
