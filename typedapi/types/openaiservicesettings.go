package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/openaisimilaritytype"
)

type OpenAIServiceSettings struct {
	ApiKey *string `json:"api_key,omitempty"`

	ClientId *string `json:"client_id,omitempty"`

	ClientSecret *string `json:"client_secret,omitempty"`

	Dimensions *int `json:"dimensions,omitempty"`

	ModelId string `json:"model_id"`

	OrganizationId *string `json:"organization_id,omitempty"`

	RateLimit *RateLimitSetting `json:"rate_limit,omitempty"`

	Scopes []string `json:"scopes,omitempty"`

	Similarity *openaisimilaritytype.OpenAISimilarityType `json:"similarity,omitempty"`

	TokenUrl *string `json:"token_url,omitempty"`

	Url *string `json:"url,omitempty"`
}

func (s *OpenAIServiceSettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewOpenAIServiceSettings() *OpenAIServiceSettings { _ = "STUB: not implemented"; return nil }

type OpenAIServiceSettingsVariant interface {
	OpenAIServiceSettingsCaster() *OpenAIServiceSettings
}

func (s *OpenAIServiceSettings) OpenAIServiceSettingsCaster() *OpenAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}
