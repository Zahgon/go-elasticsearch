package types

type ChildrenAggregate struct {
	Aggregations map[string]Aggregate `json:"-"`
	DocCount     int64                `json:"doc_count"`
	Meta         Metadata             `json:"meta,omitempty"`
}

func (s *ChildrenAggregate) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s ChildrenAggregate) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewChildrenAggregate() *ChildrenAggregate { _ = "STUB: not implemented"; return nil }
