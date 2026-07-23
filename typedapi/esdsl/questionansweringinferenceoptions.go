package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _questionAnsweringInferenceOptions struct {
	v *types.QuestionAnsweringInferenceOptions
}

func NewQuestionAnsweringInferenceOptions() *_questionAnsweringInferenceOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_questionAnsweringInferenceOptions) MaxAnswerLength(maxanswerlength int) *_questionAnsweringInferenceOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_questionAnsweringInferenceOptions) NumTopClasses(numtopclasses int) *_questionAnsweringInferenceOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_questionAnsweringInferenceOptions) ResultsField(resultsfield string) *_questionAnsweringInferenceOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_questionAnsweringInferenceOptions) Tokenization(tokenization types.TokenizationConfigContainerVariant) *_questionAnsweringInferenceOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_questionAnsweringInferenceOptions) InferenceConfigCreateContainerCaster() *types.InferenceConfigCreateContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_questionAnsweringInferenceOptions) QuestionAnsweringInferenceOptionsCaster() *types.QuestionAnsweringInferenceOptions {
	_ = "STUB: not implemented"
	return nil
}
