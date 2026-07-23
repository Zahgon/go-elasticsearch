package types

type VoyageAIServiceSettings struct {
	Dimensions *int `json:"dimensions,omitempty"`

	EmbeddingType *float32 `json:"embedding_type,omitempty"`

	ModelId string `json:"model_id"`

	RateLimit *RateLimitSetting `json:"rate_limit,omitempty"`
}

func (s *VoyageAIServiceSettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewVoyageAIServiceSettings() *VoyageAIServiceSettings { _ = "STUB: not implemented"; return nil }

type VoyageAIServiceSettingsVariant interface {
	VoyageAIServiceSettingsCaster() *VoyageAIServiceSettings
}

func (s *VoyageAIServiceSettings) VoyageAIServiceSettingsCaster() *VoyageAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}
