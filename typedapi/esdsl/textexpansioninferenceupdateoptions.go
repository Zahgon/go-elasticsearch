package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _textExpansionInferenceUpdateOptions struct {
	v *types.TextExpansionInferenceUpdateOptions
}

func NewTextExpansionInferenceUpdateOptions() *_textExpansionInferenceUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textExpansionInferenceUpdateOptions) ResultsField(resultsfield string) *_textExpansionInferenceUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textExpansionInferenceUpdateOptions) Tokenization(tokenization types.NlpTokenizationUpdateOptionsVariant) *_textExpansionInferenceUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textExpansionInferenceUpdateOptions) InferenceConfigUpdateContainerCaster() *types.InferenceConfigUpdateContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textExpansionInferenceUpdateOptions) TextExpansionInferenceUpdateOptionsCaster() *types.TextExpansionInferenceUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}
