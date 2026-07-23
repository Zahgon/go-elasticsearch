package types

type NerInferenceUpdateOptions struct {
	ResultsField *string `json:"results_field,omitempty"`

	Tokenization *NlpTokenizationUpdateOptions `json:"tokenization,omitempty"`
}

func (s *NerInferenceUpdateOptions) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewNerInferenceUpdateOptions() *NerInferenceUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}

type NerInferenceUpdateOptionsVariant interface {
	NerInferenceUpdateOptionsCaster() *NerInferenceUpdateOptions
}

func (s *NerInferenceUpdateOptions) NerInferenceUpdateOptionsCaster() *NerInferenceUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}
