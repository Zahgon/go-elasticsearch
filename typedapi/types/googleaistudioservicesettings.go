package types

type GoogleAiStudioServiceSettings struct {
	ApiKey string `json:"api_key"`

	ModelId string `json:"model_id"`

	RateLimit *RateLimitSetting `json:"rate_limit,omitempty"`
}

func (s *GoogleAiStudioServiceSettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewGoogleAiStudioServiceSettings() *GoogleAiStudioServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

type GoogleAiStudioServiceSettingsVariant interface {
	GoogleAiStudioServiceSettingsCaster() *GoogleAiStudioServiceSettings
}

func (s *GoogleAiStudioServiceSettings) GoogleAiStudioServiceSettingsCaster() *GoogleAiStudioServiceSettings {
	_ = "STUB: not implemented"
	return nil
}
