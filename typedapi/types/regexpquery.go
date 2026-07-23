package types

type RegexpQuery struct {
	Boost *float32 `json:"boost,omitempty"`

	CaseInsensitive *bool `json:"case_insensitive,omitempty"`

	Flags *string `json:"flags,omitempty"`

	MaxDeterminizedStates *int    `json:"max_determinized_states,omitempty"`
	QueryName_            *string `json:"_name,omitempty"`

	Rewrite *string `json:"rewrite,omitempty"`

	Value string `json:"value"`
}

func (s *RegexpQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRegexpQuery() *RegexpQuery { _ = "STUB: not implemented"; return nil }

type RegexpQueryVariant interface {
	RegexpQueryCaster() *RegexpQuery
}

func (s *RegexpQuery) RegexpQueryCaster() *RegexpQuery { _ = "STUB: not implemented"; return nil }
