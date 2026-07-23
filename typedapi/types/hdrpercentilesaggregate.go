package types

type HdrPercentilesAggregate struct {
	Meta   Metadata    `json:"meta,omitempty"`
	Values Percentiles `json:"values"`
}

func (s *HdrPercentilesAggregate) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewHdrPercentilesAggregate() *HdrPercentilesAggregate { _ = "STUB: not implemented"; return nil }
