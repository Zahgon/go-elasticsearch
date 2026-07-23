package types

type PercentileRanksAggregation struct {
	Field  *string `json:"field,omitempty"`
	Format *string `json:"format,omitempty"`

	Hdr *HdrMethod `json:"hdr,omitempty"`

	Keyed *bool `json:"keyed,omitempty"`

	Missing Missing `json:"missing,omitempty"`
	Script  *Script `json:"script,omitempty"`

	Tdigest *TDigest `json:"tdigest,omitempty"`

	Values *[]Float64 `json:"values,omitempty"`
}

func (s *PercentileRanksAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewPercentileRanksAggregation() *PercentileRanksAggregation {
	_ = "STUB: not implemented"
	return nil
}

type PercentileRanksAggregationVariant interface {
	PercentileRanksAggregationCaster() *PercentileRanksAggregation
}

func (s *PercentileRanksAggregation) PercentileRanksAggregationCaster() *PercentileRanksAggregation {
	_ = "STUB: not implemented"
	return nil
}
