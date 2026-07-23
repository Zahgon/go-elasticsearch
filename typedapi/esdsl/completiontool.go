package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _completionTool struct {
	v *types.CompletionTool
}

func NewCompletionTool(function types.CompletionToolFunctionVariant, type_ string) *_completionTool {
	_ = "STUB: not implemented"
	return nil
}

func (s *_completionTool) Function(function types.CompletionToolFunctionVariant) *_completionTool {
	_ = "STUB: not implemented"
	return nil
}

func (s *_completionTool) Type(type_ string) *_completionTool {
	_ = "STUB: not implemented"
	return nil
}

func (s *_completionTool) CompletionToolCaster() *types.CompletionTool {
	_ = "STUB: not implemented"
	return nil
}
