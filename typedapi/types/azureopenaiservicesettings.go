package types

type AzureOpenAIServiceSettings struct {
	ApiKey *string `json:"api_key,omitempty"`

	ApiVersion string `json:"api_version"`

	ClientId *string `json:"client_id,omitempty"`

	ClientSecret *string `json:"client_secret,omitempty"`

	DeploymentId string `json:"deployment_id"`

	EntraId *string `json:"entra_id,omitempty"`

	RateLimit *RateLimitSetting `json:"rate_limit,omitempty"`

	ResourceName string `json:"resource_name"`

	Scopes []string `json:"scopes,omitempty"`

	TenantId *string `json:"tenant_id,omitempty"`
}

func (s *AzureOpenAIServiceSettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewAzureOpenAIServiceSettings() *AzureOpenAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

type AzureOpenAIServiceSettingsVariant interface {
	AzureOpenAIServiceSettingsCaster() *AzureOpenAIServiceSettings
}

func (s *AzureOpenAIServiceSettings) AzureOpenAIServiceSettingsCaster() *AzureOpenAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}
