package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _nerInferenceUpdateOptions struct {
	v *types.NerInferenceUpdateOptions
}

func NewNerInferenceUpdateOptions() *_nerInferenceUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nerInferenceUpdateOptions) ResultsField(resultsfield string) *_nerInferenceUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nerInferenceUpdateOptions) Tokenization(tokenization types.NlpTokenizationUpdateOptionsVariant) *_nerInferenceUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nerInferenceUpdateOptions) InferenceConfigUpdateContainerCaster() *types.InferenceConfigUpdateContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nerInferenceUpdateOptions) NerInferenceUpdateOptionsCaster() *types.NerInferenceUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}
