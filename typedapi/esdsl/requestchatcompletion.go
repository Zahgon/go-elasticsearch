package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _requestChatCompletion struct {
	v *types.RequestChatCompletion
}

func NewRequestChatCompletion() *_requestChatCompletion { _ = "STUB: not implemented"; return nil }

func (s *_requestChatCompletion) MaxCompletionTokens(maxcompletiontokens int64) *_requestChatCompletion {
	_ = "STUB: not implemented"
	return nil
}

func (s *_requestChatCompletion) Messages(messages ...types.MessageVariant) *_requestChatCompletion {
	_ = "STUB: not implemented"
	return nil
}

func (s *_requestChatCompletion) MessagesValues(messagesvalues []types.Message) *_requestChatCompletion {
	_ = "STUB: not implemented"
	return nil
}

func (s *_requestChatCompletion) Model(model string) *_requestChatCompletion {
	_ = "STUB: not implemented"
	return nil
}

func (s *_requestChatCompletion) Reasoning(reasoning types.ReasoningVariant) *_requestChatCompletion {
	_ = "STUB: not implemented"
	return nil
}

func (s *_requestChatCompletion) Stop(stops ...string) *_requestChatCompletion {
	_ = "STUB: not implemented"
	return nil
}

func (s *_requestChatCompletion) Temperature(temperature float32) *_requestChatCompletion {
	_ = "STUB: not implemented"
	return nil
}

func (s *_requestChatCompletion) ToolChoice(completiontooltype types.CompletionToolTypeVariant) *_requestChatCompletion {
	_ = "STUB: not implemented"
	return nil
}

func (s *_requestChatCompletion) Tools(tools ...types.CompletionToolVariant) *_requestChatCompletion {
	_ = "STUB: not implemented"
	return nil
}

func (s *_requestChatCompletion) ToolsValues(toolsvalues []types.CompletionTool) *_requestChatCompletion {
	_ = "STUB: not implemented"
	return nil
}

func (s *_requestChatCompletion) TopP(topp float32) *_requestChatCompletion {
	_ = "STUB: not implemented"
	return nil
}

func (s *_requestChatCompletion) RequestChatCompletionCaster() *types.RequestChatCompletion {
	_ = "STUB: not implemented"
	return nil
}
