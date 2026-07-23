package types

type ContextualAIServiceSettings struct {
	ApiKey string `json:"api_key"`

	ModelId string `json:"model_id"`

	RateLimit *RateLimitSetting `json:"rate_limit,omitempty"`
}

func (s *ContextualAIServiceSettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewContextualAIServiceSettings() *ContextualAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

type ContextualAIServiceSettingsVariant interface {
	ContextualAIServiceSettingsCaster() *ContextualAIServiceSettings
}

func (s *ContextualAIServiceSettings) ContextualAIServiceSettingsCaster() *ContextualAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}
