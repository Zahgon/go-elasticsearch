package types

type Ai21ServiceSettings struct {
	ApiKey *string `json:"api_key,omitempty"`

	ModelId string `json:"model_id"`

	RateLimit *RateLimitSetting `json:"rate_limit,omitempty"`
}

func (s *Ai21ServiceSettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewAi21ServiceSettings() *Ai21ServiceSettings { _ = "STUB: not implemented"; return nil }

type Ai21ServiceSettingsVariant interface {
	Ai21ServiceSettingsCaster() *Ai21ServiceSettings
}

func (s *Ai21ServiceSettings) Ai21ServiceSettingsCaster() *Ai21ServiceSettings {
	_ = "STUB: not implemented"
	return nil
}
