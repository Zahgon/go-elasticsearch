package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _textEmbeddingInferenceOptions struct {
	v *types.TextEmbeddingInferenceOptions
}

func NewTextEmbeddingInferenceOptions() *_textEmbeddingInferenceOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textEmbeddingInferenceOptions) EmbeddingSize(embeddingsize int) *_textEmbeddingInferenceOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textEmbeddingInferenceOptions) ResultsField(resultsfield string) *_textEmbeddingInferenceOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textEmbeddingInferenceOptions) Tokenization(tokenization types.TokenizationConfigContainerVariant) *_textEmbeddingInferenceOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textEmbeddingInferenceOptions) Vocabulary(vocabulary types.VocabularyVariant) *_textEmbeddingInferenceOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textEmbeddingInferenceOptions) InferenceConfigCreateContainerCaster() *types.InferenceConfigCreateContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textEmbeddingInferenceOptions) TextEmbeddingInferenceOptionsCaster() *types.TextEmbeddingInferenceOptions {
	_ = "STUB: not implemented"
	return nil
}
