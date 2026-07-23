package types

type ParentAggregate struct {
	Aggregations map[string]Aggregate `json:"-"`
	DocCount     int64                `json:"doc_count"`
	Meta         Metadata             `json:"meta,omitempty"`
}

func (s *ParentAggregate) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s ParentAggregate) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewParentAggregate() *ParentAggregate { _ = "STUB: not implemented"; return nil }
