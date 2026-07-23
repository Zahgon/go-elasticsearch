package types

type OpenAITaskSettings struct {
	Headers map[string]string `json:"headers,omitempty"`

	User *string `json:"user,omitempty"`
}

func (s *OpenAITaskSettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewOpenAITaskSettings() *OpenAITaskSettings { _ = "STUB: not implemented"; return nil }

type OpenAITaskSettingsVariant interface {
	OpenAITaskSettingsCaster() *OpenAITaskSettings
}

func (s *OpenAITaskSettings) OpenAITaskSettingsCaster() *OpenAITaskSettings {
	_ = "STUB: not implemented"
	return nil
}
