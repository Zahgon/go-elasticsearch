package types

type ToolCall struct {
	Function ToolCallFunction `json:"function"`

	Id string `json:"id"`

	Type string `json:"type"`
}

func (s *ToolCall) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewToolCall() *ToolCall { _ = "STUB: not implemented"; return nil }

type ToolCallVariant interface {
	ToolCallCaster() *ToolCall
}

func (s *ToolCall) ToolCallCaster() *ToolCall { _ = "STUB: not implemented"; return nil }
