package types

type TDigestPercentilesAggregate struct {
	Meta   Metadata    `json:"meta,omitempty"`
	Values Percentiles `json:"values"`
}

func (s *TDigestPercentilesAggregate) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewTDigestPercentilesAggregate() *TDigestPercentilesAggregate {
	_ = "STUB: not implemented"
	return nil
}
