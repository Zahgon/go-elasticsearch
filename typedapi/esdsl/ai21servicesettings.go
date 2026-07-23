package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _ai21ServiceSettings struct {
	v *types.Ai21ServiceSettings
}

func NewAi21ServiceSettings(modelid string) *_ai21ServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ai21ServiceSettings) ApiKey(apikey string) *_ai21ServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ai21ServiceSettings) ModelId(modelid string) *_ai21ServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ai21ServiceSettings) RateLimit(ratelimit types.RateLimitSettingVariant) *_ai21ServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ai21ServiceSettings) Ai21ServiceSettingsCaster() *types.Ai21ServiceSettings {
	_ = "STUB: not implemented"
	return nil
}
