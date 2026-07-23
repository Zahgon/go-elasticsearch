package types

type HasParentQuery struct {
	Boost *float32 `json:"boost,omitempty"`

	IgnoreUnmapped *bool `json:"ignore_unmapped,omitempty"`

	InnerHits *InnerHits `json:"inner_hits,omitempty"`

	ParentType string `json:"parent_type"`

	Query      Query   `json:"query"`
	QueryName_ *string `json:"_name,omitempty"`

	Score *bool `json:"score,omitempty"`
}

func (s *HasParentQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewHasParentQuery() *HasParentQuery { _ = "STUB: not implemented"; return nil }

type HasParentQueryVariant interface {
	HasParentQueryCaster() *HasParentQuery
}

func (s *HasParentQuery) HasParentQueryCaster() *HasParentQuery {
	_ = "STUB: not implemented"
	return nil
}
