package types

type FireworksAITaskSettings struct {
	Headers map[string]string `json:"headers,omitempty"`

	User *string `json:"user,omitempty"`
}

func (s *FireworksAITaskSettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewFireworksAITaskSettings() *FireworksAITaskSettings { _ = "STUB: not implemented"; return nil }

type FireworksAITaskSettingsVariant interface {
	FireworksAITaskSettingsCaster() *FireworksAITaskSettings
}

func (s *FireworksAITaskSettings) FireworksAITaskSettingsCaster() *FireworksAITaskSettings {
	_ = "STUB: not implemented"
	return nil
}
