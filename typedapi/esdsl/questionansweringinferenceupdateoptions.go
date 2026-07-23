package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _questionAnsweringInferenceUpdateOptions struct {
	v *types.QuestionAnsweringInferenceUpdateOptions
}

func NewQuestionAnsweringInferenceUpdateOptions(question string) *_questionAnsweringInferenceUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_questionAnsweringInferenceUpdateOptions) MaxAnswerLength(maxanswerlength int) *_questionAnsweringInferenceUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_questionAnsweringInferenceUpdateOptions) NumTopClasses(numtopclasses int) *_questionAnsweringInferenceUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_questionAnsweringInferenceUpdateOptions) Question(question string) *_questionAnsweringInferenceUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_questionAnsweringInferenceUpdateOptions) ResultsField(resultsfield string) *_questionAnsweringInferenceUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_questionAnsweringInferenceUpdateOptions) Tokenization(tokenization types.NlpTokenizationUpdateOptionsVariant) *_questionAnsweringInferenceUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_questionAnsweringInferenceUpdateOptions) InferenceConfigUpdateContainerCaster() *types.InferenceConfigUpdateContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_questionAnsweringInferenceUpdateOptions) QuestionAnsweringInferenceUpdateOptionsCaster() *types.QuestionAnsweringInferenceUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}
