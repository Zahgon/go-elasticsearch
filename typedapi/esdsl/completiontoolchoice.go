package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _completionToolChoice struct {
	v *types.CompletionToolChoice
}

func NewCompletionToolChoice(function types.CompletionToolChoiceFunctionVariant, type_ string) *_completionToolChoice {
	_ = "STUB: not implemented"
	return nil
}

func (s *_completionToolChoice) Function(function types.CompletionToolChoiceFunctionVariant) *_completionToolChoice {
	_ = "STUB: not implemented"
	return nil
}

func (s *_completionToolChoice) Type(type_ string) *_completionToolChoice {
	_ = "STUB: not implemented"
	return nil
}

func (s *_completionToolChoice) CompletionToolChoiceCaster() *types.CompletionToolChoice {
	_ = "STUB: not implemented"
	return nil
}
