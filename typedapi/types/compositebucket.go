package types

type CompositeBucket struct {
	Aggregations map[string]Aggregate  `json:"-"`
	DocCount     int64                 `json:"doc_count"`
	Key          CompositeAggregateKey `json:"key"`
}

func (s *CompositeBucket) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s CompositeBucket) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewCompositeBucket() *CompositeBucket { _ = "STUB: not implemented"; return nil }
