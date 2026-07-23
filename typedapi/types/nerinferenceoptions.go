package types

type NerInferenceOptions struct {
	ClassificationLabels []string `json:"classification_labels,omitempty"`

	ResultsField *string `json:"results_field,omitempty"`

	Tokenization *TokenizationConfigContainer `json:"tokenization,omitempty"`
	Vocabulary   *Vocabulary                  `json:"vocabulary,omitempty"`
}

func (s *NerInferenceOptions) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewNerInferenceOptions() *NerInferenceOptions { _ = "STUB: not implemented"; return nil }

type NerInferenceOptionsVariant interface {
	NerInferenceOptionsCaster() *NerInferenceOptions
}

func (s *NerInferenceOptions) NerInferenceOptionsCaster() *NerInferenceOptions {
	_ = "STUB: not implemented"
	return nil
}
