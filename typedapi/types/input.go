package types

type Input struct {
	FieldNames []string `json:"field_names"`
}

func (s *Input) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewInput() *Input { _ = "STUB: not implemented"; return nil }

type InputVariant interface {
	InputCaster() *Input
}

func (s *Input) InputCaster() *Input { _ = "STUB: not implemented"; return nil }
