package types

type AzureAiStudioTaskSettings struct {
	DoSample *float32 `json:"do_sample,omitempty"`

	MaxNewTokens *int `json:"max_new_tokens,omitempty"`

	ReturnDocuments *bool `json:"return_documents,omitempty"`

	Temperature *float32 `json:"temperature,omitempty"`

	TopN *int `json:"top_n,omitempty"`

	TopP *float32 `json:"top_p,omitempty"`

	User *string `json:"user,omitempty"`
}

func (s *AzureAiStudioTaskSettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewAzureAiStudioTaskSettings() *AzureAiStudioTaskSettings {
	_ = "STUB: not implemented"
	return nil
}

type AzureAiStudioTaskSettingsVariant interface {
	AzureAiStudioTaskSettingsCaster() *AzureAiStudioTaskSettings
}

func (s *AzureAiStudioTaskSettings) AzureAiStudioTaskSettingsCaster() *AzureAiStudioTaskSettings {
	_ = "STUB: not implemented"
	return nil
}
