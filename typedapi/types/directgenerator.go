package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/suggestmode"
)

type DirectGenerator struct {
	Field string `json:"field"`

	MaxEdits *int `json:"max_edits,omitempty"`

	MaxInspections *float32 `json:"max_inspections,omitempty"`

	MaxTermFreq *float32 `json:"max_term_freq,omitempty"`

	MinDocFreq *float32 `json:"min_doc_freq,omitempty"`

	MinWordLength *int `json:"min_word_length,omitempty"`

	PostFilter *string `json:"post_filter,omitempty"`

	PreFilter *string `json:"pre_filter,omitempty"`

	PrefixLength *int `json:"prefix_length,omitempty"`

	Size *int `json:"size,omitempty"`

	SuggestMode *suggestmode.SuggestMode `json:"suggest_mode,omitempty"`
}

func (s *DirectGenerator) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewDirectGenerator() *DirectGenerator { _ = "STUB: not implemented"; return nil }

type DirectGeneratorVariant interface {
	DirectGeneratorCaster() *DirectGenerator
}

func (s *DirectGenerator) DirectGeneratorCaster() *DirectGenerator {
	_ = "STUB: not implemented"
	return nil
}
