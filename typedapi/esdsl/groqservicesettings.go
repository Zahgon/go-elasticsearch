package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _groqServiceSettings struct {
	v *types.GroqServiceSettings
}

func NewGroqServiceSettings(modelid string) *_groqServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_groqServiceSettings) ApiKey(apikey string) *_groqServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_groqServiceSettings) ModelId(modelid string) *_groqServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_groqServiceSettings) RateLimit(ratelimit types.RateLimitSettingVariant) *_groqServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_groqServiceSettings) GroqServiceSettingsCaster() *types.GroqServiceSettings {
	_ = "STUB: not implemented"
	return nil
}
