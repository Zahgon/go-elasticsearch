package types

type GlobalAggregate struct {
	Aggregations map[string]Aggregate `json:"-"`
	DocCount     int64                `json:"doc_count"`
	Meta         Metadata             `json:"meta,omitempty"`
}

func (s *GlobalAggregate) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s GlobalAggregate) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewGlobalAggregate() *GlobalAggregate { _ = "STUB: not implemented"; return nil }
