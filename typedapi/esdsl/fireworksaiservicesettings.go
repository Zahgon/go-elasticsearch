package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/fireworksaisimilaritytype"
)

type _fireworksAIServiceSettings struct {
	v *types.FireworksAIServiceSettings
}

func NewFireworksAIServiceSettings(apikey string, modelid string) *_fireworksAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fireworksAIServiceSettings) ApiKey(apikey string) *_fireworksAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fireworksAIServiceSettings) Dimensions(dimensions int) *_fireworksAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fireworksAIServiceSettings) ModelId(modelid string) *_fireworksAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fireworksAIServiceSettings) RateLimit(ratelimit types.RateLimitSettingVariant) *_fireworksAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fireworksAIServiceSettings) Similarity(similarity fireworksaisimilaritytype.FireworksAISimilarityType) *_fireworksAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fireworksAIServiceSettings) Url(url string) *_fireworksAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fireworksAIServiceSettings) FireworksAIServiceSettingsCaster() *types.FireworksAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}
