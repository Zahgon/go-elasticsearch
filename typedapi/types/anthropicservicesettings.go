package types

type AnthropicServiceSettings struct {
	ApiKey string `json:"api_key"`

	ModelId string `json:"model_id"`

	RateLimit *RateLimitSetting `json:"rate_limit,omitempty"`
}

func (s *AnthropicServiceSettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewAnthropicServiceSettings() *AnthropicServiceSettings { _ = "STUB: not implemented"; return nil }

type AnthropicServiceSettingsVariant interface {
	AnthropicServiceSettingsCaster() *AnthropicServiceSettings
}

func (s *AnthropicServiceSettings) AnthropicServiceSettingsCaster() *AnthropicServiceSettings {
	_ = "STUB: not implemented"
	return nil
}
