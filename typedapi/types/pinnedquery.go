package types

type PinnedQuery struct {
	Boost *float32 `json:"boost,omitempty"`

	Docs []PinnedDoc `json:"docs,omitempty"`

	Ids []string `json:"ids,omitempty"`

	Organic    Query   `json:"organic"`
	QueryName_ *string `json:"_name,omitempty"`
}

func (s *PinnedQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewPinnedQuery() *PinnedQuery { _ = "STUB: not implemented"; return nil }

type PinnedQueryVariant interface {
	PinnedQueryCaster() *PinnedQuery
}

func (s *PinnedQuery) PinnedQueryCaster() *PinnedQuery { _ = "STUB: not implemented"; return nil }
