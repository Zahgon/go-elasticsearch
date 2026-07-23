package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _anthropicServiceSettings struct {
	v *types.AnthropicServiceSettings
}

func NewAnthropicServiceSettings(apikey string, modelid string) *_anthropicServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_anthropicServiceSettings) ApiKey(apikey string) *_anthropicServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_anthropicServiceSettings) ModelId(modelid string) *_anthropicServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_anthropicServiceSettings) RateLimit(ratelimit types.RateLimitSettingVariant) *_anthropicServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_anthropicServiceSettings) AnthropicServiceSettingsCaster() *types.AnthropicServiceSettings {
	_ = "STUB: not implemented"
	return nil
}
