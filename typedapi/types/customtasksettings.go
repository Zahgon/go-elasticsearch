package types

type CustomTaskSettings struct {
	Parameters map[string]CustomTaskParameter `json:"parameters,omitempty"`
}

func NewCustomTaskSettings() *CustomTaskSettings { _ = "STUB: not implemented"; return nil }

type CustomTaskSettingsVariant interface {
	CustomTaskSettingsCaster() *CustomTaskSettings
}

func (s *CustomTaskSettings) CustomTaskSettingsCaster() *CustomTaskSettings {
	_ = "STUB: not implemented"
	return nil
}
