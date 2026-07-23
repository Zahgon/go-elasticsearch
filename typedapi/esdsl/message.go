package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _message struct {
	v *types.Message
}

func NewMessage(role string) *_message { _ = "STUB: not implemented"; return nil }

func (s *_message) Content(messagecontent types.MessageContentVariant) *_message {
	_ = "STUB: not implemented"
	return nil
}

func (s *_message) Reasoning(reasoning string) *_message { _ = "STUB: not implemented"; return nil }

func (s *_message) ReasoningDetails(reasoningdetails ...types.ReasoningDetailVariant) *_message {
	_ = "STUB: not implemented"
	return nil
}

func (s *_message) ReasoningDetailsValues(reasoningdetailsvalues []types.ReasoningDetail) *_message {
	_ = "STUB: not implemented"
	return nil
}

func (s *_message) Role(role string) *_message { _ = "STUB: not implemented"; return nil }

func (s *_message) ToolCallId(id string) *_message { _ = "STUB: not implemented"; return nil }

func (s *_message) ToolCalls(toolcalls ...types.ToolCallVariant) *_message {
	_ = "STUB: not implemented"
	return nil
}

func (s *_message) ToolCallsValues(toolcallsvalues []types.ToolCall) *_message {
	_ = "STUB: not implemented"
	return nil
}

func (s *_message) MessageCaster() *types.Message { _ = "STUB: not implemented"; return nil }
