package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/tokenchar"
)

type NGramTokenizer struct {
	CustomTokenChars *string               `json:"custom_token_chars,omitempty"`
	MaxGram          *int                  `json:"max_gram,omitempty"`
	MinGram          *int                  `json:"min_gram,omitempty"`
	TokenChars       []tokenchar.TokenChar `json:"token_chars,omitempty"`
	Type             string                `json:"type,omitempty"`
	Version          *string               `json:"version,omitempty"`
}

func (s *NGramTokenizer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s NGramTokenizer) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewNGramTokenizer() *NGramTokenizer { _ = "STUB: not implemented"; return nil }

type NGramTokenizerVariant interface {
	NGramTokenizerCaster() *NGramTokenizer
}

func (s *NGramTokenizer) NGramTokenizerCaster() *NGramTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *NGramTokenizer) TokenizerDefinitionCaster() *TokenizerDefinition {
	_ = "STUB: not implemented"
	return nil
}
