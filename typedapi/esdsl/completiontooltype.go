package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _completionToolType struct {
	v types.CompletionToolType
}

func NewCompletionToolType() *_completionToolType { _ = "STUB: not implemented"; return nil }

func (u *_completionToolType) String(string string) *_completionToolType {
	_ = "STUB: not implemented"
	return nil
}

func (u *_completionToolType) CompletionToolChoice(completiontoolchoice types.CompletionToolChoiceVariant) *_completionToolType {
	_ = "STUB: not implemented"
	return nil
}

func (u *_completionToolChoice) CompletionToolTypeCaster() *types.CompletionToolType {
	_ = "STUB: not implemented"
	return nil
}

func (u *_completionToolType) CompletionToolTypeCaster() *types.CompletionToolType {
	_ = "STUB: not implemented"
	return nil
}
