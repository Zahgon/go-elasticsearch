package types

type ScriptQuery struct {
	Boost      *float32 `json:"boost,omitempty"`
	QueryName_ *string  `json:"_name,omitempty"`

	Script Script `json:"script"`
}

func (s *ScriptQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewScriptQuery() *ScriptQuery { _ = "STUB: not implemented"; return nil }

type ScriptQueryVariant interface {
	ScriptQueryCaster() *ScriptQuery
}

func (s *ScriptQuery) ScriptQueryCaster() *ScriptQuery { _ = "STUB: not implemented"; return nil }
