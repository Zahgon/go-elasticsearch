package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/googlemodelgardenprovider"
)

type GoogleVertexAIServiceSettings struct {
	Dimensions *int `json:"dimensions,omitempty"`

	Location *string `json:"location,omitempty"`

	MaxBatchSize *int `json:"max_batch_size,omitempty"`

	ModelId *string `json:"model_id,omitempty"`

	ProjectId *string `json:"project_id,omitempty"`

	Provider *googlemodelgardenprovider.GoogleModelGardenProvider `json:"provider,omitempty"`

	RateLimit *RateLimitSetting `json:"rate_limit,omitempty"`

	ServiceAccountJson string `json:"service_account_json"`

	StreamingUrl *string `json:"streaming_url,omitempty"`

	Url *string `json:"url,omitempty"`
}

func (s *GoogleVertexAIServiceSettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewGoogleVertexAIServiceSettings() *GoogleVertexAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

type GoogleVertexAIServiceSettingsVariant interface {
	GoogleVertexAIServiceSettingsCaster() *GoogleVertexAIServiceSettings
}

func (s *GoogleVertexAIServiceSettings) GoogleVertexAIServiceSettingsCaster() *GoogleVertexAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}
