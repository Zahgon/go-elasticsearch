package types

type PercentilesBucketAggregate struct {
	Meta   Metadata    `json:"meta,omitempty"`
	Values Percentiles `json:"values"`
}

func (s *PercentilesBucketAggregate) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewPercentilesBucketAggregate() *PercentilesBucketAggregate {
	_ = "STUB: not implemented"
	return nil
}
