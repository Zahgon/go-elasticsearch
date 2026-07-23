package types

type AzureAiStudioServiceSettings struct {
	ApiKey string `json:"api_key"`

	EndpointType string `json:"endpoint_type"`

	Provider string `json:"provider"`

	RateLimit *RateLimitSetting `json:"rate_limit,omitempty"`

	Target string `json:"target"`
}

func (s *AzureAiStudioServiceSettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewAzureAiStudioServiceSettings() *AzureAiStudioServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

type AzureAiStudioServiceSettingsVariant interface {
	AzureAiStudioServiceSettingsCaster() *AzureAiStudioServiceSettings
}

func (s *AzureAiStudioServiceSettings) AzureAiStudioServiceSettingsCaster() *AzureAiStudioServiceSettings {
	_ = "STUB: not implemented"
	return nil
}
