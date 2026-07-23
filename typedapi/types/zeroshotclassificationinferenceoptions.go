package types

type ZeroShotClassificationInferenceOptions struct {
	ClassificationLabels []string `json:"classification_labels"`

	HypothesisTemplate *string `json:"hypothesis_template,omitempty"`

	Labels []string `json:"labels,omitempty"`

	MultiLabel *bool `json:"multi_label,omitempty"`

	ResultsField *string `json:"results_field,omitempty"`

	Tokenization *TokenizationConfigContainer `json:"tokenization,omitempty"`
}

func (s *ZeroShotClassificationInferenceOptions) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewZeroShotClassificationInferenceOptions() *ZeroShotClassificationInferenceOptions {
	_ = "STUB: not implemented"
	return nil
}

type ZeroShotClassificationInferenceOptionsVariant interface {
	ZeroShotClassificationInferenceOptionsCaster() *ZeroShotClassificationInferenceOptions
}

func (s *ZeroShotClassificationInferenceOptions) ZeroShotClassificationInferenceOptionsCaster() *ZeroShotClassificationInferenceOptions {
	_ = "STUB: not implemented"
	return nil
}
