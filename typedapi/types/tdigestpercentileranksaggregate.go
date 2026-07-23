package types

type TDigestPercentileRanksAggregate struct {
	Meta   Metadata    `json:"meta,omitempty"`
	Values Percentiles `json:"values"`
}

func (s *TDigestPercentileRanksAggregate) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewTDigestPercentileRanksAggregate() *TDigestPercentileRanksAggregate {
	_ = "STUB: not implemented"
	return nil
}
