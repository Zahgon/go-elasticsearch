package types

type TopMetricsAggregate struct {
	Meta Metadata     `json:"meta,omitempty"`
	Top  []TopMetrics `json:"top"`
}

func (s *TopMetricsAggregate) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewTopMetricsAggregate() *TopMetricsAggregate { _ = "STUB: not implemented"; return nil }
