package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _contextualAIServiceSettings struct {
	v *types.ContextualAIServiceSettings
}

func NewContextualAIServiceSettings(apikey string, modelid string) *_contextualAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_contextualAIServiceSettings) ApiKey(apikey string) *_contextualAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_contextualAIServiceSettings) ModelId(modelid string) *_contextualAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_contextualAIServiceSettings) RateLimit(ratelimit types.RateLimitSettingVariant) *_contextualAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_contextualAIServiceSettings) ContextualAIServiceSettingsCaster() *types.ContextualAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}
