package types

type ElisionTokenFilter struct {
	Articles []string `json:"articles,omitempty"`

	ArticlesCase Stringifiedboolean `json:"articles_case,omitempty"`

	ArticlesPath *string `json:"articles_path,omitempty"`
	Type         string  `json:"type,omitempty"`
	Version      *string `json:"version,omitempty"`
}

func (s *ElisionTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s ElisionTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewElisionTokenFilter() *ElisionTokenFilter { _ = "STUB: not implemented"; return nil }

type ElisionTokenFilterVariant interface {
	ElisionTokenFilterCaster() *ElisionTokenFilter
}

func (s *ElisionTokenFilter) ElisionTokenFilterCaster() *ElisionTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *ElisionTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
