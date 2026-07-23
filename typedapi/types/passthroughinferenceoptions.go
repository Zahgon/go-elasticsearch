package types

type PassThroughInferenceOptions struct {
	ResultsField *string `json:"results_field,omitempty"`

	Tokenization *TokenizationConfigContainer `json:"tokenization,omitempty"`
	Vocabulary   *Vocabulary                  `json:"vocabulary,omitempty"`
}

func (s *PassThroughInferenceOptions) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewPassThroughInferenceOptions() *PassThroughInferenceOptions {
	_ = "STUB: not implemented"
	return nil
}

type PassThroughInferenceOptionsVariant interface {
	PassThroughInferenceOptionsCaster() *PassThroughInferenceOptions
}

func (s *PassThroughInferenceOptions) PassThroughInferenceOptionsCaster() *PassThroughInferenceOptions {
	_ = "STUB: not implemented"
	return nil
}
