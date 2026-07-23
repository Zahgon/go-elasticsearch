package types

type TextExpansionInferenceOptions struct {
	ResultsField *string `json:"results_field,omitempty"`

	Tokenization *TokenizationConfigContainer `json:"tokenization,omitempty"`
	Vocabulary   *Vocabulary                  `json:"vocabulary,omitempty"`
}

func (s *TextExpansionInferenceOptions) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewTextExpansionInferenceOptions() *TextExpansionInferenceOptions {
	_ = "STUB: not implemented"
	return nil
}

type TextExpansionInferenceOptionsVariant interface {
	TextExpansionInferenceOptionsCaster() *TextExpansionInferenceOptions
}

func (s *TextExpansionInferenceOptions) TextExpansionInferenceOptionsCaster() *TextExpansionInferenceOptions {
	_ = "STUB: not implemented"
	return nil
}
