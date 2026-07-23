package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/edgengramside"
)

type EdgeNGramTokenFilter struct {
	MaxGram *int `json:"max_gram,omitempty"`

	MinGram *int `json:"min_gram,omitempty"`

	PreserveOriginal Stringifiedboolean `json:"preserve_original,omitempty"`

	Side    *edgengramside.EdgeNGramSide `json:"side,omitempty"`
	Type    string                       `json:"type,omitempty"`
	Version *string                      `json:"version,omitempty"`
}

func (s *EdgeNGramTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s EdgeNGramTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewEdgeNGramTokenFilter() *EdgeNGramTokenFilter { _ = "STUB: not implemented"; return nil }

type EdgeNGramTokenFilterVariant interface {
	EdgeNGramTokenFilterCaster() *EdgeNGramTokenFilter
}

func (s *EdgeNGramTokenFilter) EdgeNGramTokenFilterCaster() *EdgeNGramTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *EdgeNGramTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
