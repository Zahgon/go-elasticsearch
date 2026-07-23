package types

type CompositeAggregate struct {
	AfterKey CompositeAggregateKey  `json:"after_key,omitempty"`
	Buckets  BucketsCompositeBucket `json:"buckets"`
	Meta     Metadata               `json:"meta,omitempty"`
}

func (s *CompositeAggregate) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewCompositeAggregate() *CompositeAggregate { _ = "STUB: not implemented"; return nil }
