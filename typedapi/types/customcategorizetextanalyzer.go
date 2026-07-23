package types

type CustomCategorizeTextAnalyzer struct {
	CharFilter []string `json:"char_filter,omitempty"`
	Filter     []string `json:"filter,omitempty"`
	Tokenizer  *string  `json:"tokenizer,omitempty"`
}

func (s *CustomCategorizeTextAnalyzer) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewCustomCategorizeTextAnalyzer() *CustomCategorizeTextAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

type CustomCategorizeTextAnalyzerVariant interface {
	CustomCategorizeTextAnalyzerCaster() *CustomCategorizeTextAnalyzer
}

func (s *CustomCategorizeTextAnalyzer) CustomCategorizeTextAnalyzerCaster() *CustomCategorizeTextAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *CustomCategorizeTextAnalyzer) CategorizeTextAnalyzerCaster() *CategorizeTextAnalyzer {
	_ = "STUB: not implemented"
	return nil
}
