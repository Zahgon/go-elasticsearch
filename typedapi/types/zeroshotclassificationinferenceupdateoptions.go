package types

type ZeroShotClassificationInferenceUpdateOptions struct {
	Labels []string `json:"labels"`

	MultiLabel *bool `json:"multi_label,omitempty"`

	ResultsField *string `json:"results_field,omitempty"`

	Tokenization *NlpTokenizationUpdateOptions `json:"tokenization,omitempty"`
}

func (s *ZeroShotClassificationInferenceUpdateOptions) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewZeroShotClassificationInferenceUpdateOptions() *ZeroShotClassificationInferenceUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}

type ZeroShotClassificationInferenceUpdateOptionsVariant interface {
	ZeroShotClassificationInferenceUpdateOptionsCaster() *ZeroShotClassificationInferenceUpdateOptions
}

func (s *ZeroShotClassificationInferenceUpdateOptions) ZeroShotClassificationInferenceUpdateOptionsCaster() *ZeroShotClassificationInferenceUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}
