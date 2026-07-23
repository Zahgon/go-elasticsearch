package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _zeroShotClassificationInferenceUpdateOptions struct {
	v *types.ZeroShotClassificationInferenceUpdateOptions
}

func NewZeroShotClassificationInferenceUpdateOptions() *_zeroShotClassificationInferenceUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_zeroShotClassificationInferenceUpdateOptions) Labels(labels ...string) *_zeroShotClassificationInferenceUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_zeroShotClassificationInferenceUpdateOptions) MultiLabel(multilabel bool) *_zeroShotClassificationInferenceUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_zeroShotClassificationInferenceUpdateOptions) ResultsField(resultsfield string) *_zeroShotClassificationInferenceUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_zeroShotClassificationInferenceUpdateOptions) Tokenization(tokenization types.NlpTokenizationUpdateOptionsVariant) *_zeroShotClassificationInferenceUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_zeroShotClassificationInferenceUpdateOptions) InferenceConfigUpdateContainerCaster() *types.InferenceConfigUpdateContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_zeroShotClassificationInferenceUpdateOptions) ZeroShotClassificationInferenceUpdateOptionsCaster() *types.ZeroShotClassificationInferenceUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}
