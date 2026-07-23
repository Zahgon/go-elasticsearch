package types

type DelayedDataCheckConfig struct {
	CheckWindow Duration `json:"check_window,omitempty"`

	Enabled bool `json:"enabled"`
}

func (s *DelayedDataCheckConfig) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDelayedDataCheckConfig() *DelayedDataCheckConfig { _ = "STUB: not implemented"; return nil }

type DelayedDataCheckConfigVariant interface {
	DelayedDataCheckConfigCaster() *DelayedDataCheckConfig
}

func (s *DelayedDataCheckConfig) DelayedDataCheckConfigCaster() *DelayedDataCheckConfig {
	_ = "STUB: not implemented"
	return nil
}
