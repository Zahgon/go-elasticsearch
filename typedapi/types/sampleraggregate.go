package types

type SamplerAggregate struct {
	Aggregations map[string]Aggregate `json:"-"`
	DocCount     int64                `json:"doc_count"`
	Meta         Metadata             `json:"meta,omitempty"`
}

func (s *SamplerAggregate) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s SamplerAggregate) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewSamplerAggregate() *SamplerAggregate { _ = "STUB: not implemented"; return nil }
