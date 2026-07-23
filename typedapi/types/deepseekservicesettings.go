package types

type DeepSeekServiceSettings struct {
	ApiKey string `json:"api_key"`

	ModelId string `json:"model_id"`

	Url *string `json:"url,omitempty"`
}

func (s *DeepSeekServiceSettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDeepSeekServiceSettings() *DeepSeekServiceSettings { _ = "STUB: not implemented"; return nil }

type DeepSeekServiceSettingsVariant interface {
	DeepSeekServiceSettingsCaster() *DeepSeekServiceSettings
}

func (s *DeepSeekServiceSettings) DeepSeekServiceSettingsCaster() *DeepSeekServiceSettings {
	_ = "STUB: not implemented"
	return nil
}
