package types

type TextEmbeddingInferenceUpdateOptions struct {
	ResultsField *string                       `json:"results_field,omitempty"`
	Tokenization *NlpTokenizationUpdateOptions `json:"tokenization,omitempty"`
}

func (s *TextEmbeddingInferenceUpdateOptions) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewTextEmbeddingInferenceUpdateOptions() *TextEmbeddingInferenceUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}

type TextEmbeddingInferenceUpdateOptionsVariant interface {
	TextEmbeddingInferenceUpdateOptionsCaster() *TextEmbeddingInferenceUpdateOptions
}

func (s *TextEmbeddingInferenceUpdateOptions) TextEmbeddingInferenceUpdateOptionsCaster() *TextEmbeddingInferenceUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}
