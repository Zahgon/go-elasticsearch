package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/cohereembeddingtype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/coheresimilaritytype"
)

type CohereServiceSettings struct {
	ApiKey string `json:"api_key"`

	EmbeddingType *cohereembeddingtype.CohereEmbeddingType `json:"embedding_type,omitempty"`

	ModelId string `json:"model_id"`

	RateLimit *RateLimitSetting `json:"rate_limit,omitempty"`

	Similarity *coheresimilaritytype.CohereSimilarityType `json:"similarity,omitempty"`
}

func (s *CohereServiceSettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewCohereServiceSettings() *CohereServiceSettings { _ = "STUB: not implemented"; return nil }

type CohereServiceSettingsVariant interface {
	CohereServiceSettingsCaster() *CohereServiceSettings
}

func (s *CohereServiceSettings) CohereServiceSettingsCaster() *CohereServiceSettings {
	_ = "STUB: not implemented"
	return nil
}
