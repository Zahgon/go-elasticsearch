package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/nvidiasimilaritytype"
)

type NvidiaServiceSettings struct {
	ApiKey string `json:"api_key"`

	MaxInputTokens *int `json:"max_input_tokens,omitempty"`

	ModelId string `json:"model_id"`

	RateLimit *RateLimitSetting `json:"rate_limit,omitempty"`

	Similarity *nvidiasimilaritytype.NvidiaSimilarityType `json:"similarity,omitempty"`

	Url *string `json:"url,omitempty"`
}

func (s *NvidiaServiceSettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewNvidiaServiceSettings() *NvidiaServiceSettings { _ = "STUB: not implemented"; return nil }

type NvidiaServiceSettingsVariant interface {
	NvidiaServiceSettingsCaster() *NvidiaServiceSettings
}

func (s *NvidiaServiceSettings) NvidiaServiceSettingsCaster() *NvidiaServiceSettings {
	_ = "STUB: not implemented"
	return nil
}
