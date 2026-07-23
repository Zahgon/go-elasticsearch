package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/fireworksaisimilaritytype"
)

type FireworksAIServiceSettings struct {
	ApiKey string `json:"api_key"`

	Dimensions *int `json:"dimensions,omitempty"`

	ModelId string `json:"model_id"`

	RateLimit *RateLimitSetting `json:"rate_limit,omitempty"`

	Similarity *fireworksaisimilaritytype.FireworksAISimilarityType `json:"similarity,omitempty"`

	Url *string `json:"url,omitempty"`
}

func (s *FireworksAIServiceSettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewFireworksAIServiceSettings() *FireworksAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

type FireworksAIServiceSettingsVariant interface {
	FireworksAIServiceSettingsCaster() *FireworksAIServiceSettings
}

func (s *FireworksAIServiceSettings) FireworksAIServiceSettingsCaster() *FireworksAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}
