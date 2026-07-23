package types

type ContextMethod struct {
	Name       string               `json:"name"`
	Params     []ContextMethodParam `json:"params"`
	ReturnType string               `json:"return_type"`
}

func (s *ContextMethod) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewContextMethod() *ContextMethod { _ = "STUB: not implemented"; return nil }
