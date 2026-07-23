package types

type NestedAggregate struct {
	Aggregations map[string]Aggregate `json:"-"`
	DocCount     int64                `json:"doc_count"`
	Meta         Metadata             `json:"meta,omitempty"`
}

func (s *NestedAggregate) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s NestedAggregate) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewNestedAggregate() *NestedAggregate { _ = "STUB: not implemented"; return nil }
