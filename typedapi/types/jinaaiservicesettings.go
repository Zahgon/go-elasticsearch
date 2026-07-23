package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/jinaaielementtype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/jinaaisimilaritytype"
)

type JinaAIServiceSettings struct {
	ApiKey string `json:"api_key"`

	Dimensions *int `json:"dimensions,omitempty"`

	EmbeddingType *jinaaielementtype.JinaAIElementType `json:"embedding_type,omitempty"`

	ModelId string `json:"model_id"`

	MultimodalModel *bool `json:"multimodal_model,omitempty"`

	RateLimit *RateLimitSetting `json:"rate_limit,omitempty"`

	Similarity *jinaaisimilaritytype.JinaAISimilarityType `json:"similarity,omitempty"`
}

func (s *JinaAIServiceSettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewJinaAIServiceSettings() *JinaAIServiceSettings { _ = "STUB: not implemented"; return nil }

type JinaAIServiceSettingsVariant interface {
	JinaAIServiceSettingsCaster() *JinaAIServiceSettings
}

func (s *JinaAIServiceSettings) JinaAIServiceSettingsCaster() *JinaAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}
