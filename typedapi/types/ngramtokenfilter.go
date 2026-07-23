package types

type NGramTokenFilter struct {
	MaxGram *int `json:"max_gram,omitempty"`

	MinGram *int `json:"min_gram,omitempty"`

	PreserveOriginal Stringifiedboolean `json:"preserve_original,omitempty"`
	Type             string             `json:"type,omitempty"`
	Version          *string            `json:"version,omitempty"`
}

func (s *NGramTokenFilter) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s NGramTokenFilter) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewNGramTokenFilter() *NGramTokenFilter { _ = "STUB: not implemented"; return nil }

type NGramTokenFilterVariant interface {
	NGramTokenFilterCaster() *NGramTokenFilter
}

func (s *NGramTokenFilter) NGramTokenFilterCaster() *NGramTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *NGramTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
