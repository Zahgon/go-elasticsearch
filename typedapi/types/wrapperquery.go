package types

type WrapperQuery struct {
	Boost *float32 `json:"boost,omitempty"`

	Query      string  `json:"query"`
	QueryName_ *string `json:"_name,omitempty"`
}

func (s *WrapperQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewWrapperQuery() *WrapperQuery { _ = "STUB: not implemented"; return nil }

type WrapperQueryVariant interface {
	WrapperQueryCaster() *WrapperQuery
}

func (s *WrapperQuery) WrapperQueryCaster() *WrapperQuery { _ = "STUB: not implemented"; return nil }
