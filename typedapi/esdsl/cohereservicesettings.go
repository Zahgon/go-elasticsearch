package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/cohereembeddingtype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/coheresimilaritytype"
)

type _cohereServiceSettings struct {
	v *types.CohereServiceSettings
}

func NewCohereServiceSettings(apikey string, modelid string) *_cohereServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_cohereServiceSettings) ApiKey(apikey string) *_cohereServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_cohereServiceSettings) EmbeddingType(embeddingtype cohereembeddingtype.CohereEmbeddingType) *_cohereServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_cohereServiceSettings) ModelId(modelid string) *_cohereServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_cohereServiceSettings) RateLimit(ratelimit types.RateLimitSettingVariant) *_cohereServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_cohereServiceSettings) Similarity(similarity coheresimilaritytype.CohereSimilarityType) *_cohereServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_cohereServiceSettings) CohereServiceSettingsCaster() *types.CohereServiceSettings {
	_ = "STUB: not implemented"
	return nil
}
