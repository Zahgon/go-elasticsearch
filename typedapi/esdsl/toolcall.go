package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _toolCall struct {
	v *types.ToolCall
}

func NewToolCall(function types.ToolCallFunctionVariant, type_ string) *_toolCall {
	_ = "STUB: not implemented"
	return nil
}

func (s *_toolCall) Function(function types.ToolCallFunctionVariant) *_toolCall {
	_ = "STUB: not implemented"
	return nil
}

func (s *_toolCall) Id(id string) *_toolCall { _ = "STUB: not implemented"; return nil }

func (s *_toolCall) Type(type_ string) *_toolCall { _ = "STUB: not implemented"; return nil }

func (s *_toolCall) ToolCallCaster() *types.ToolCall { _ = "STUB: not implemented"; return nil }
