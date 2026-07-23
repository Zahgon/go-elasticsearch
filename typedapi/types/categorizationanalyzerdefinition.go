package types

type CategorizationAnalyzerDefinition struct {
	CharFilter []CharFilter `json:"char_filter,omitempty"`

	Filter []TokenFilter `json:"filter,omitempty"`

	Tokenizer Tokenizer `json:"tokenizer,omitempty"`
}

func (s *CategorizationAnalyzerDefinition) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewCategorizationAnalyzerDefinition() *CategorizationAnalyzerDefinition {
	_ = "STUB: not implemented"
	return nil
}

type CategorizationAnalyzerDefinitionVariant interface {
	CategorizationAnalyzerDefinitionCaster() *CategorizationAnalyzerDefinition
}

func (s *CategorizationAnalyzerDefinition) CategorizationAnalyzerDefinitionCaster() *CategorizationAnalyzerDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (s *CategorizationAnalyzerDefinition) CategorizationAnalyzerCaster() *CategorizationAnalyzer {
	_ = "STUB: not implemented"
	return nil
}
