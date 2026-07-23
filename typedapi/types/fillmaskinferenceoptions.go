package types

type FillMaskInferenceOptions struct {
	MaskToken *string `json:"mask_token,omitempty"`

	NumTopClasses *int `json:"num_top_classes,omitempty"`

	ResultsField *string `json:"results_field,omitempty"`

	Tokenization *TokenizationConfigContainer `json:"tokenization,omitempty"`
	Vocabulary   *Vocabulary                  `json:"vocabulary,omitempty"`
}

func (s *FillMaskInferenceOptions) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewFillMaskInferenceOptions() *FillMaskInferenceOptions { _ = "STUB: not implemented"; return nil }

type FillMaskInferenceOptionsVariant interface {
	FillMaskInferenceOptionsCaster() *FillMaskInferenceOptions
}

func (s *FillMaskInferenceOptions) FillMaskInferenceOptionsCaster() *FillMaskInferenceOptions {
	_ = "STUB: not implemented"
	return nil
}
