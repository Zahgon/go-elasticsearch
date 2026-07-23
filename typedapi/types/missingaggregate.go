package types

type MissingAggregate struct {
	Aggregations map[string]Aggregate `json:"-"`
	DocCount     int64                `json:"doc_count"`
	Meta         Metadata             `json:"meta,omitempty"`
}

func (s *MissingAggregate) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s MissingAggregate) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewMissingAggregate() *MissingAggregate { _ = "STUB: not implemented"; return nil }
