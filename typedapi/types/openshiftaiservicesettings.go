package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/openshiftaisimilaritytype"
)

type OpenShiftAiServiceSettings struct {
	ApiKey string `json:"api_key"`

	MaxInputTokens *int `json:"max_input_tokens,omitempty"`

	ModelId *string `json:"model_id,omitempty"`

	RateLimit *RateLimitSetting `json:"rate_limit,omitempty"`

	Similarity *openshiftaisimilaritytype.OpenShiftAiSimilarityType `json:"similarity,omitempty"`

	Url string `json:"url"`
}

func (s *OpenShiftAiServiceSettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewOpenShiftAiServiceSettings() *OpenShiftAiServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

type OpenShiftAiServiceSettingsVariant interface {
	OpenShiftAiServiceSettingsCaster() *OpenShiftAiServiceSettings
}

func (s *OpenShiftAiServiceSettings) OpenShiftAiServiceSettingsCaster() *OpenShiftAiServiceSettings {
	_ = "STUB: not implemented"
	return nil
}
