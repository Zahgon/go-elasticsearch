package types

type SlowlogTresholds struct {
	Fetch *SlowlogTresholdLevels `json:"fetch,omitempty"`
	Query *SlowlogTresholdLevels `json:"query,omitempty"`
}

func NewSlowlogTresholds() *SlowlogTresholds { _ = "STUB: not implemented"; return nil }

type SlowlogTresholdsVariant interface {
	SlowlogTresholdsCaster() *SlowlogTresholds
}

func (s *SlowlogTresholds) SlowlogTresholdsCaster() *SlowlogTresholds {
	_ = "STUB: not implemented"
	return nil
}
