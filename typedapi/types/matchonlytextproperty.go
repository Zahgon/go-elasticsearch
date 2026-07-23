package types

type MatchOnlyTextProperty struct {
	CopyTo []string `json:"copy_to,omitempty"`

	Fields map[string]Property `json:"fields,omitempty"`

	Meta map[string]string `json:"meta,omitempty"`
	Type string            `json:"type,omitempty"`
}

func (s *MatchOnlyTextProperty) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s MatchOnlyTextProperty) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewMatchOnlyTextProperty() *MatchOnlyTextProperty { _ = "STUB: not implemented"; return nil }

type MatchOnlyTextPropertyVariant interface {
	MatchOnlyTextPropertyCaster() *MatchOnlyTextProperty
}

func (s *MatchOnlyTextProperty) MatchOnlyTextPropertyCaster() *MatchOnlyTextProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *MatchOnlyTextProperty) PropertyCaster() *Property { _ = "STUB: not implemented"; return nil }
