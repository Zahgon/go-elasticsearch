package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _textEmbeddingInferenceUpdateOptions struct {
	v *types.TextEmbeddingInferenceUpdateOptions
}

func NewTextEmbeddingInferenceUpdateOptions() *_textEmbeddingInferenceUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textEmbeddingInferenceUpdateOptions) ResultsField(resultsfield string) *_textEmbeddingInferenceUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textEmbeddingInferenceUpdateOptions) Tokenization(tokenization types.NlpTokenizationUpdateOptionsVariant) *_textEmbeddingInferenceUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textEmbeddingInferenceUpdateOptions) InferenceConfigUpdateContainerCaster() *types.InferenceConfigUpdateContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textEmbeddingInferenceUpdateOptions) TextEmbeddingInferenceUpdateOptionsCaster() *types.TextEmbeddingInferenceUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}
