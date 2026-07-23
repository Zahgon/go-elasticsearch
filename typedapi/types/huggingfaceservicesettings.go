package types

type HuggingFaceServiceSettings struct {
	ApiKey string `json:"api_key"`

	ModelId *string `json:"model_id,omitempty"`

	RateLimit *RateLimitSetting `json:"rate_limit,omitempty"`

	Url string `json:"url"`
}

func (s *HuggingFaceServiceSettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewHuggingFaceServiceSettings() *HuggingFaceServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

type HuggingFaceServiceSettingsVariant interface {
	HuggingFaceServiceSettingsCaster() *HuggingFaceServiceSettings
}

func (s *HuggingFaceServiceSettings) HuggingFaceServiceSettingsCaster() *HuggingFaceServiceSettings {
	_ = "STUB: not implemented"
	return nil
}
