package types

type SlowlogTresholdLevels struct {
	Debug Duration `json:"debug,omitempty"`
	Info  Duration `json:"info,omitempty"`
	Trace Duration `json:"trace,omitempty"`
	Warn  Duration `json:"warn,omitempty"`
}

func (s *SlowlogTresholdLevels) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewSlowlogTresholdLevels() *SlowlogTresholdLevels { _ = "STUB: not implemented"; return nil }

type SlowlogTresholdLevelsVariant interface {
	SlowlogTresholdLevelsCaster() *SlowlogTresholdLevels
}

func (s *SlowlogTresholdLevels) SlowlogTresholdLevelsCaster() *SlowlogTresholdLevels {
	_ = "STUB: not implemented"
	return nil
}
