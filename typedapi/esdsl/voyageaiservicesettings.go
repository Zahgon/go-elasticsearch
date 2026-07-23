package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _voyageAIServiceSettings struct {
	v *types.VoyageAIServiceSettings
}

func NewVoyageAIServiceSettings(modelid string) *_voyageAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_voyageAIServiceSettings) Dimensions(dimensions int) *_voyageAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_voyageAIServiceSettings) EmbeddingType(embeddingtype float32) *_voyageAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_voyageAIServiceSettings) ModelId(modelid string) *_voyageAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_voyageAIServiceSettings) RateLimit(ratelimit types.RateLimitSettingVariant) *_voyageAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_voyageAIServiceSettings) VoyageAIServiceSettingsCaster() *types.VoyageAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}
