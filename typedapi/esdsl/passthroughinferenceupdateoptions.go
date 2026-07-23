package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _passThroughInferenceUpdateOptions struct {
	v *types.PassThroughInferenceUpdateOptions
}

func NewPassThroughInferenceUpdateOptions() *_passThroughInferenceUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_passThroughInferenceUpdateOptions) ResultsField(resultsfield string) *_passThroughInferenceUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_passThroughInferenceUpdateOptions) Tokenization(tokenization types.NlpTokenizationUpdateOptionsVariant) *_passThroughInferenceUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_passThroughInferenceUpdateOptions) InferenceConfigUpdateContainerCaster() *types.InferenceConfigUpdateContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_passThroughInferenceUpdateOptions) PassThroughInferenceUpdateOptionsCaster() *types.PassThroughInferenceUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}
