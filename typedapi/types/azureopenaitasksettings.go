package types

type AzureOpenAITaskSettings struct {
	Headers map[string]string `json:"headers,omitempty"`

	User *string `json:"user,omitempty"`
}

func (s *AzureOpenAITaskSettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewAzureOpenAITaskSettings() *AzureOpenAITaskSettings { _ = "STUB: not implemented"; return nil }

type AzureOpenAITaskSettingsVariant interface {
	AzureOpenAITaskSettingsCaster() *AzureOpenAITaskSettings
}

func (s *AzureOpenAITaskSettings) AzureOpenAITaskSettingsCaster() *AzureOpenAITaskSettings {
	_ = "STUB: not implemented"
	return nil
}
