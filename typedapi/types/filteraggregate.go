package types

type FilterAggregate struct {
	Aggregations map[string]Aggregate `json:"-"`
	DocCount     int64                `json:"doc_count"`
	Meta         Metadata             `json:"meta,omitempty"`
}

func (s *FilterAggregate) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s FilterAggregate) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewFilterAggregate() *FilterAggregate { _ = "STUB: not implemented"; return nil }
