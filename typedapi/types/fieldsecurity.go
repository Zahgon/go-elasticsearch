package types

type FieldSecurity struct {
	Except []string `json:"except,omitempty"`
	Grant  []string `json:"grant,omitempty"`
}

func (s *FieldSecurity) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewFieldSecurity() *FieldSecurity { _ = "STUB: not implemented"; return nil }

type FieldSecurityVariant interface {
	FieldSecurityCaster() *FieldSecurity
}

func (s *FieldSecurity) FieldSecurityCaster() *FieldSecurity { _ = "STUB: not implemented"; return nil }
