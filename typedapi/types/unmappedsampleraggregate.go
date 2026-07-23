package types

type UnmappedSamplerAggregate struct {
	Aggregations map[string]Aggregate `json:"-"`
	DocCount     int64                `json:"doc_count"`
	Meta         Metadata             `json:"meta,omitempty"`
}

func (s *UnmappedSamplerAggregate) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s UnmappedSamplerAggregate) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewUnmappedSamplerAggregate() *UnmappedSamplerAggregate { _ = "STUB: not implemented"; return nil }
