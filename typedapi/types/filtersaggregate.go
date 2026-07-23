package types

type FiltersAggregate struct {
	Buckets BucketsFiltersBucket `json:"buckets"`
	Meta    Metadata             `json:"meta,omitempty"`
}

func (s *FiltersAggregate) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewFiltersAggregate() *FiltersAggregate { _ = "STUB: not implemented"; return nil }
