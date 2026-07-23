package types

type TextClassificationInferenceOptions struct {
	ClassificationLabels []string `json:"classification_labels,omitempty"`

	NumTopClasses *int `json:"num_top_classes,omitempty"`

	ResultsField *string `json:"results_field,omitempty"`

	Tokenization *TokenizationConfigContainer `json:"tokenization,omitempty"`
	Vocabulary   *Vocabulary                  `json:"vocabulary,omitempty"`
}

func (s *TextClassificationInferenceOptions) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewTextClassificationInferenceOptions() *TextClassificationInferenceOptions {
	_ = "STUB: not implemented"
	return nil
}

type TextClassificationInferenceOptionsVariant interface {
	TextClassificationInferenceOptionsCaster() *TextClassificationInferenceOptions
}

func (s *TextClassificationInferenceOptions) TextClassificationInferenceOptionsCaster() *TextClassificationInferenceOptions {
	_ = "STUB: not implemented"
	return nil
}
