package types

type HdrPercentileRanksAggregate struct {
	Meta   Metadata    `json:"meta,omitempty"`
	Values Percentiles `json:"values"`
}

func (s *HdrPercentileRanksAggregate) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewHdrPercentileRanksAggregate() *HdrPercentileRanksAggregate {
	_ = "STUB: not implemented"
	return nil
}
