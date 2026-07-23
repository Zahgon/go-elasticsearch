package types

type TextClassificationInferenceUpdateOptions struct {
	ClassificationLabels []string `json:"classification_labels,omitempty"`

	NumTopClasses *int `json:"num_top_classes,omitempty"`

	ResultsField *string `json:"results_field,omitempty"`

	Tokenization *NlpTokenizationUpdateOptions `json:"tokenization,omitempty"`
}

func (s *TextClassificationInferenceUpdateOptions) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewTextClassificationInferenceUpdateOptions() *TextClassificationInferenceUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}

type TextClassificationInferenceUpdateOptionsVariant interface {
	TextClassificationInferenceUpdateOptionsCaster() *TextClassificationInferenceUpdateOptions
}

func (s *TextClassificationInferenceUpdateOptions) TextClassificationInferenceUpdateOptionsCaster() *TextClassificationInferenceUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}
