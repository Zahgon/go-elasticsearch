package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _nerInferenceOptions struct {
	v *types.NerInferenceOptions
}

func NewNerInferenceOptions() *_nerInferenceOptions { _ = "STUB: not implemented"; return nil }

func (s *_nerInferenceOptions) ClassificationLabels(classificationlabels ...string) *_nerInferenceOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nerInferenceOptions) ResultsField(resultsfield string) *_nerInferenceOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nerInferenceOptions) Tokenization(tokenization types.TokenizationConfigContainerVariant) *_nerInferenceOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nerInferenceOptions) Vocabulary(vocabulary types.VocabularyVariant) *_nerInferenceOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nerInferenceOptions) InferenceConfigCreateContainerCaster() *types.InferenceConfigCreateContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nerInferenceOptions) NerInferenceOptionsCaster() *types.NerInferenceOptions {
	_ = "STUB: not implemented"
	return nil
}
