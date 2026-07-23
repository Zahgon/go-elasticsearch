package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _textExpansionInferenceOptions struct {
	v *types.TextExpansionInferenceOptions
}

func NewTextExpansionInferenceOptions() *_textExpansionInferenceOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textExpansionInferenceOptions) ResultsField(resultsfield string) *_textExpansionInferenceOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textExpansionInferenceOptions) Tokenization(tokenization types.TokenizationConfigContainerVariant) *_textExpansionInferenceOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textExpansionInferenceOptions) Vocabulary(vocabulary types.VocabularyVariant) *_textExpansionInferenceOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textExpansionInferenceOptions) InferenceConfigCreateContainerCaster() *types.InferenceConfigCreateContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textExpansionInferenceOptions) TextExpansionInferenceOptionsCaster() *types.TextExpansionInferenceOptions {
	_ = "STUB: not implemented"
	return nil
}
