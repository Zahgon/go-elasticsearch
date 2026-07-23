package types

type TextExpansionInferenceUpdateOptions struct {
	ResultsField *string                       `json:"results_field,omitempty"`
	Tokenization *NlpTokenizationUpdateOptions `json:"tokenization,omitempty"`
}

func (s *TextExpansionInferenceUpdateOptions) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewTextExpansionInferenceUpdateOptions() *TextExpansionInferenceUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}

type TextExpansionInferenceUpdateOptionsVariant interface {
	TextExpansionInferenceUpdateOptionsCaster() *TextExpansionInferenceUpdateOptions
}

func (s *TextExpansionInferenceUpdateOptions) TextExpansionInferenceUpdateOptionsCaster() *TextExpansionInferenceUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}
