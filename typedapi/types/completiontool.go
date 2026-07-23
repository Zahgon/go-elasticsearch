package types

type CompletionTool struct {
	Function CompletionToolFunction `json:"function"`

	Type string `json:"type"`
}

func (s *CompletionTool) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewCompletionTool() *CompletionTool { _ = "STUB: not implemented"; return nil }

type CompletionToolVariant interface {
	CompletionToolCaster() *CompletionTool
}

func (s *CompletionTool) CompletionToolCaster() *CompletionTool {
	_ = "STUB: not implemented"
	return nil
}
