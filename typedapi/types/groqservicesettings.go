package types

type GroqServiceSettings struct {
	ApiKey *string `json:"api_key,omitempty"`

	ModelId string `json:"model_id"`

	RateLimit *RateLimitSetting `json:"rate_limit,omitempty"`
}

func (s *GroqServiceSettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewGroqServiceSettings() *GroqServiceSettings { _ = "STUB: not implemented"; return nil }

type GroqServiceSettingsVariant interface {
	GroqServiceSettingsCaster() *GroqServiceSettings
}

func (s *GroqServiceSettings) GroqServiceSettingsCaster() *GroqServiceSettings {
	_ = "STUB: not implemented"
	return nil
}
