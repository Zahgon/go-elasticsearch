package types

type RequestChatCompletion struct {
	MaxCompletionTokens *int64 `json:"max_completion_tokens,omitempty"`

	Messages []Message `json:"messages"`

	Model *string `json:"model,omitempty"`

	Reasoning *Reasoning `json:"reasoning,omitempty"`

	Stop []string `json:"stop,omitempty"`

	Temperature *float32 `json:"temperature,omitempty"`

	ToolChoice CompletionToolType `json:"tool_choice,omitempty"`

	Tools []CompletionTool `json:"tools,omitempty"`

	TopP *float32 `json:"top_p,omitempty"`
}

func (s *RequestChatCompletion) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewRequestChatCompletion() *RequestChatCompletion { _ = "STUB: not implemented"; return nil }

type RequestChatCompletionVariant interface {
	RequestChatCompletionCaster() *RequestChatCompletion
}

func (s *RequestChatCompletion) RequestChatCompletionCaster() *RequestChatCompletion {
	_ = "STUB: not implemented"
	return nil
}
