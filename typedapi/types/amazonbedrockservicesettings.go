package types

type AmazonBedrockServiceSettings struct {
	AccessKey string `json:"access_key"`

	Model string `json:"model"`

	Provider *string `json:"provider,omitempty"`

	RateLimit *RateLimitSetting `json:"rate_limit,omitempty"`

	Region string `json:"region"`

	SecretKey string `json:"secret_key"`
}

func (s *AmazonBedrockServiceSettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewAmazonBedrockServiceSettings() *AmazonBedrockServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

type AmazonBedrockServiceSettingsVariant interface {
	AmazonBedrockServiceSettingsCaster() *AmazonBedrockServiceSettings
}

func (s *AmazonBedrockServiceSettings) AmazonBedrockServiceSettingsCaster() *AmazonBedrockServiceSettings {
	_ = "STUB: not implemented"
	return nil
}
