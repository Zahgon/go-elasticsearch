package types

type SamplerAggregation struct {
	ShardSize *int `json:"shard_size,omitempty"`
}

func (s *SamplerAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewSamplerAggregation() *SamplerAggregation { _ = "STUB: not implemented"; return nil }

type SamplerAggregationVariant interface {
	SamplerAggregationCaster() *SamplerAggregation
}

func (s *SamplerAggregation) SamplerAggregationCaster() *SamplerAggregation {
	_ = "STUB: not implemented"
	return nil
}
