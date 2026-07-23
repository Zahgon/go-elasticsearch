package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/llamasimilaritytype"
)

type LlamaServiceSettings struct {
	MaxInputTokens *int `json:"max_input_tokens,omitempty"`

	ModelId string `json:"model_id"`

	RateLimit *RateLimitSetting `json:"rate_limit,omitempty"`

	Similarity *llamasimilaritytype.LlamaSimilarityType `json:"similarity,omitempty"`

	Url string `json:"url"`
}

func (s *LlamaServiceSettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewLlamaServiceSettings() *LlamaServiceSettings { _ = "STUB: not implemented"; return nil }

type LlamaServiceSettingsVariant interface {
	LlamaServiceSettingsCaster() *LlamaServiceSettings
}

func (s *LlamaServiceSettings) LlamaServiceSettingsCaster() *LlamaServiceSettings {
	_ = "STUB: not implemented"
	return nil
}
