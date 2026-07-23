package types

type PassThroughInferenceUpdateOptions struct {
	ResultsField *string `json:"results_field,omitempty"`

	Tokenization *NlpTokenizationUpdateOptions `json:"tokenization,omitempty"`
}

func (s *PassThroughInferenceUpdateOptions) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewPassThroughInferenceUpdateOptions() *PassThroughInferenceUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}

type PassThroughInferenceUpdateOptionsVariant interface {
	PassThroughInferenceUpdateOptionsCaster() *PassThroughInferenceUpdateOptions
}

func (s *PassThroughInferenceUpdateOptions) PassThroughInferenceUpdateOptionsCaster() *PassThroughInferenceUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}
