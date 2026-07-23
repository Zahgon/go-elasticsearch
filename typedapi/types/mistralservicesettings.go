package types

type MistralServiceSettings struct {
	ApiKey string `json:"api_key"`

	MaxInputTokens *int `json:"max_input_tokens,omitempty"`

	Model string `json:"model"`

	RateLimit *RateLimitSetting `json:"rate_limit,omitempty"`
}

func (s *MistralServiceSettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewMistralServiceSettings() *MistralServiceSettings { _ = "STUB: not implemented"; return nil }

type MistralServiceSettingsVariant interface {
	MistralServiceSettingsCaster() *MistralServiceSettings
}

func (s *MistralServiceSettings) MistralServiceSettingsCaster() *MistralServiceSettings {
	_ = "STUB: not implemented"
	return nil
}
