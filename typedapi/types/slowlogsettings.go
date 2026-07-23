package types

type SlowlogSettings struct {
	Level     *string           `json:"level,omitempty"`
	Reformat  *bool             `json:"reformat,omitempty"`
	Source    *int              `json:"source,omitempty"`
	Threshold *SlowlogTresholds `json:"threshold,omitempty"`
}

func (s *SlowlogSettings) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSlowlogSettings() *SlowlogSettings { _ = "STUB: not implemented"; return nil }

type SlowlogSettingsVariant interface {
	SlowlogSettingsCaster() *SlowlogSettings
}

func (s *SlowlogSettings) SlowlogSettingsCaster() *SlowlogSettings {
	_ = "STUB: not implemented"
	return nil
}
