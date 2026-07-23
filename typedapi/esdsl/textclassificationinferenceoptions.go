package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _textClassificationInferenceOptions struct {
	v *types.TextClassificationInferenceOptions
}

func NewTextClassificationInferenceOptions() *_textClassificationInferenceOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textClassificationInferenceOptions) ClassificationLabels(classificationlabels ...string) *_textClassificationInferenceOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textClassificationInferenceOptions) NumTopClasses(numtopclasses int) *_textClassificationInferenceOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textClassificationInferenceOptions) ResultsField(resultsfield string) *_textClassificationInferenceOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textClassificationInferenceOptions) Tokenization(tokenization types.TokenizationConfigContainerVariant) *_textClassificationInferenceOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textClassificationInferenceOptions) Vocabulary(vocabulary types.VocabularyVariant) *_textClassificationInferenceOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textClassificationInferenceOptions) InferenceConfigCreateContainerCaster() *types.InferenceConfigCreateContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textClassificationInferenceOptions) TextClassificationInferenceOptionsCaster() *types.TextClassificationInferenceOptions {
	_ = "STUB: not implemented"
	return nil
}
