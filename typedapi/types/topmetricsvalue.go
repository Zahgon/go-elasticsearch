package types

type TopMetricsValue struct {
	Field string `json:"field"`
}

func (s *TopMetricsValue) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewTopMetricsValue() *TopMetricsValue { _ = "STUB: not implemented"; return nil }

type TopMetricsValueVariant interface {
	TopMetricsValueCaster() *TopMetricsValue
}

func (s *TopMetricsValue) TopMetricsValueCaster() *TopMetricsValue {
	_ = "STUB: not implemented"
	return nil
}
