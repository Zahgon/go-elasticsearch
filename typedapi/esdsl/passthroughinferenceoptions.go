package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _passThroughInferenceOptions struct {
	v *types.PassThroughInferenceOptions
}

func NewPassThroughInferenceOptions() *_passThroughInferenceOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_passThroughInferenceOptions) ResultsField(resultsfield string) *_passThroughInferenceOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_passThroughInferenceOptions) Tokenization(tokenization types.TokenizationConfigContainerVariant) *_passThroughInferenceOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_passThroughInferenceOptions) Vocabulary(vocabulary types.VocabularyVariant) *_passThroughInferenceOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_passThroughInferenceOptions) InferenceConfigCreateContainerCaster() *types.InferenceConfigCreateContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_passThroughInferenceOptions) PassThroughInferenceOptionsCaster() *types.PassThroughInferenceOptions {
	_ = "STUB: not implemented"
	return nil
}
