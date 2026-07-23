package types

type PercentilesAggregation struct {
	Field  *string `json:"field,omitempty"`
	Format *string `json:"format,omitempty"`

	Hdr *HdrMethod `json:"hdr,omitempty"`

	Keyed *bool `json:"keyed,omitempty"`

	Missing Missing `json:"missing,omitempty"`

	Percents []Float64 `json:"percents,omitempty"`
	Script   *Script   `json:"script,omitempty"`

	Tdigest *TDigest `json:"tdigest,omitempty"`
}

func (s *PercentilesAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewPercentilesAggregation() *PercentilesAggregation { _ = "STUB: not implemented"; return nil }

type PercentilesAggregationVariant interface {
	PercentilesAggregationCaster() *PercentilesAggregation
}

func (s *PercentilesAggregation) PercentilesAggregationCaster() *PercentilesAggregation {
	_ = "STUB: not implemented"
	return nil
}
