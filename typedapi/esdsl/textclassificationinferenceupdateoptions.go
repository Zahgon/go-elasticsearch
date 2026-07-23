package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _textClassificationInferenceUpdateOptions struct {
	v *types.TextClassificationInferenceUpdateOptions
}

func NewTextClassificationInferenceUpdateOptions() *_textClassificationInferenceUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textClassificationInferenceUpdateOptions) ClassificationLabels(classificationlabels ...string) *_textClassificationInferenceUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textClassificationInferenceUpdateOptions) NumTopClasses(numtopclasses int) *_textClassificationInferenceUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textClassificationInferenceUpdateOptions) ResultsField(resultsfield string) *_textClassificationInferenceUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textClassificationInferenceUpdateOptions) Tokenization(tokenization types.NlpTokenizationUpdateOptionsVariant) *_textClassificationInferenceUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textClassificationInferenceUpdateOptions) InferenceConfigUpdateContainerCaster() *types.InferenceConfigUpdateContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textClassificationInferenceUpdateOptions) TextClassificationInferenceUpdateOptionsCaster() *types.TextClassificationInferenceUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}
