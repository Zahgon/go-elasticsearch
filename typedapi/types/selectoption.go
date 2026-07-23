package types

type SelectOption struct {
	Label string      `json:"label"`
	Value ScalarValue `json:"value"`
}

func (s *SelectOption) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSelectOption() *SelectOption { _ = "STUB: not implemented"; return nil }

type SelectOptionVariant interface {
	SelectOptionCaster() *SelectOption
}

func (s *SelectOption) SelectOptionCaster() *SelectOption { _ = "STUB: not implemented"; return nil }
