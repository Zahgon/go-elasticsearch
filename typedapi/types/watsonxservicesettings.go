package types

type WatsonxServiceSettings struct {
	ApiKey string `json:"api_key"`

	ApiVersion string `json:"api_version"`

	ModelId string `json:"model_id"`

	ProjectId string `json:"project_id"`

	RateLimit *RateLimitSetting `json:"rate_limit,omitempty"`

	Url string `json:"url"`
}

func (s *WatsonxServiceSettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewWatsonxServiceSettings() *WatsonxServiceSettings { _ = "STUB: not implemented"; return nil }

type WatsonxServiceSettingsVariant interface {
	WatsonxServiceSettingsCaster() *WatsonxServiceSettings
}

func (s *WatsonxServiceSettings) WatsonxServiceSettingsCaster() *WatsonxServiceSettings {
	_ = "STUB: not implemented"
	return nil
}
