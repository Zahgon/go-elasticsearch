package types

type ContextMethodParam struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func (s *ContextMethodParam) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewContextMethodParam() *ContextMethodParam { _ = "STUB: not implemented"; return nil }
