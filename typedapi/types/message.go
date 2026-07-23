package types

type Message struct {
	Content MessageContent `json:"content,omitempty"`

	Reasoning *string `json:"reasoning,omitempty"`

	ReasoningDetails []ReasoningDetail `json:"reasoning_details,omitempty"`

	Role string `json:"role"`

	ToolCallId *string `json:"tool_call_id,omitempty"`

	ToolCalls []ToolCall `json:"tool_calls,omitempty"`
}

func (s *Message) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewMessage() *Message { _ = "STUB: not implemented"; return nil }

type MessageVariant interface {
	MessageCaster() *Message
}

func (s *Message) MessageCaster() *Message { _ = "STUB: not implemented"; return nil }
