package types

type FillMaskInferenceUpdateOptions struct {
	NumTopClasses *int `json:"num_top_classes,omitempty"`

	ResultsField *string `json:"results_field,omitempty"`

	Tokenization *NlpTokenizationUpdateOptions `json:"tokenization,omitempty"`
}

func (s *FillMaskInferenceUpdateOptions) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewFillMaskInferenceUpdateOptions() *FillMaskInferenceUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}

type FillMaskInferenceUpdateOptionsVariant interface {
	FillMaskInferenceUpdateOptionsCaster() *FillMaskInferenceUpdateOptions
}

func (s *FillMaskInferenceUpdateOptions) FillMaskInferenceUpdateOptionsCaster() *FillMaskInferenceUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}
