package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/stringdistance"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/suggestmode"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/suggestsort"
)

type TermSuggester struct {
	Analyzer *string `json:"analyzer,omitempty"`

	Field          string `json:"field"`
	LowercaseTerms *bool  `json:"lowercase_terms,omitempty"`

	MaxEdits *int `json:"max_edits,omitempty"`

	MaxInspections *int `json:"max_inspections,omitempty"`

	MaxTermFreq *float32 `json:"max_term_freq,omitempty"`

	MinDocFreq *float32 `json:"min_doc_freq,omitempty"`

	MinWordLength *int `json:"min_word_length,omitempty"`

	PrefixLength *int `json:"prefix_length,omitempty"`

	ShardSize *int `json:"shard_size,omitempty"`

	Size *int `json:"size,omitempty"`

	Sort *suggestsort.SuggestSort `json:"sort,omitempty"`

	StringDistance *stringdistance.StringDistance `json:"string_distance,omitempty"`

	SuggestMode *suggestmode.SuggestMode `json:"suggest_mode,omitempty"`
}

func (s *TermSuggester) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewTermSuggester() *TermSuggester { _ = "STUB: not implemented"; return nil }

type TermSuggesterVariant interface {
	TermSuggesterCaster() *TermSuggester
}

func (s *TermSuggester) TermSuggesterCaster() *TermSuggester { _ = "STUB: not implemented"; return nil }
