package types

type RandomSamplerAggregation struct {
	Probability Float64 `json:"probability"`

	Seed *int `json:"seed,omitempty"`

	ShardSeed *int `json:"shard_seed,omitempty"`
}

func (s *RandomSamplerAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewRandomSamplerAggregation() *RandomSamplerAggregation { _ = "STUB: not implemented"; return nil }

type RandomSamplerAggregationVariant interface {
	RandomSamplerAggregationCaster() *RandomSamplerAggregation
}

func (s *RandomSamplerAggregation) RandomSamplerAggregationCaster() *RandomSamplerAggregation {
	_ = "STUB: not implemented"
	return nil
}
