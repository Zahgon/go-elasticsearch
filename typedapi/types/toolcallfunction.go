package types

type ToolCallFunction struct {
	Arguments string `json:"arguments"`

	Name string `json:"name"`
}

func (s *ToolCallFunction) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewToolCallFunction() *ToolCallFunction { _ = "STUB: not implemented"; return nil }

type ToolCallFunctionVariant interface {
	ToolCallFunctionCaster() *ToolCallFunction
}

func (s *ToolCallFunction) ToolCallFunctionCaster() *ToolCallFunction {
	_ = "STUB: not implemented"
	return nil
}
