package types

type PatternCaptureTokenFilter struct {
	Patterns []string `json:"patterns"`

	PreserveOriginal Stringifiedboolean `json:"preserve_original,omitempty"`
	Type             string             `json:"type,omitempty"`
	Version          *string            `json:"version,omitempty"`
}

func (s *PatternCaptureTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s PatternCaptureTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewPatternCaptureTokenFilter() *PatternCaptureTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

type PatternCaptureTokenFilterVariant interface {
	PatternCaptureTokenFilterCaster() *PatternCaptureTokenFilter
}

func (s *PatternCaptureTokenFilter) PatternCaptureTokenFilterCaster() *PatternCaptureTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *PatternCaptureTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
