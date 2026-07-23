package types

type PhraseSuggester struct {
	Analyzer *string `json:"analyzer,omitempty"`

	Collate *PhraseSuggestCollate `json:"collate,omitempty"`

	Confidence *Float64 `json:"confidence,omitempty"`

	DirectGenerator []DirectGenerator `json:"direct_generator,omitempty"`

	Field         string `json:"field"`
	ForceUnigrams *bool  `json:"force_unigrams,omitempty"`

	GramSize *int `json:"gram_size,omitempty"`

	Highlight *PhraseSuggestHighlight `json:"highlight,omitempty"`

	MaxErrors *Float64 `json:"max_errors,omitempty"`

	RealWordErrorLikelihood *Float64 `json:"real_word_error_likelihood,omitempty"`

	Separator *string `json:"separator,omitempty"`

	ShardSize *int `json:"shard_size,omitempty"`

	Size *int `json:"size,omitempty"`

	Smoothing  *SmoothingModelContainer `json:"smoothing,omitempty"`
	TokenLimit *int                     `json:"token_limit,omitempty"`
}

func (s *PhraseSuggester) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewPhraseSuggester() *PhraseSuggester { _ = "STUB: not implemented"; return nil }

type PhraseSuggesterVariant interface {
	PhraseSuggesterCaster() *PhraseSuggester
}

func (s *PhraseSuggester) PhraseSuggesterCaster() *PhraseSuggester {
	_ = "STUB: not implemented"
	return nil
}
