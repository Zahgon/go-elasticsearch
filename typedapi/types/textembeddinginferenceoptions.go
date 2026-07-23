package types

type TextEmbeddingInferenceOptions struct {
	EmbeddingSize *int `json:"embedding_size,omitempty"`

	ResultsField *string `json:"results_field,omitempty"`

	Tokenization *TokenizationConfigContainer `json:"tokenization,omitempty"`
	Vocabulary   *Vocabulary                  `json:"vocabulary,omitempty"`
}

func (s *TextEmbeddingInferenceOptions) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewTextEmbeddingInferenceOptions() *TextEmbeddingInferenceOptions {
	_ = "STUB: not implemented"
	return nil
}

type TextEmbeddingInferenceOptionsVariant interface {
	TextEmbeddingInferenceOptionsCaster() *TextEmbeddingInferenceOptions
}

func (s *TextEmbeddingInferenceOptions) TextEmbeddingInferenceOptionsCaster() *TextEmbeddingInferenceOptions {
	_ = "STUB: not implemented"
	return nil
}
