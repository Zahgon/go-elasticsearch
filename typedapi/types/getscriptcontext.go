package types

type GetScriptContext struct {
	Methods []ContextMethod `json:"methods"`
	Name    string          `json:"name"`
}

func (s *GetScriptContext) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewGetScriptContext() *GetScriptContext { _ = "STUB: not implemented"; return nil }
