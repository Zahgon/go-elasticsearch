package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _mistralServiceSettings struct {
	v *types.MistralServiceSettings
}

func NewMistralServiceSettings(apikey string, model string) *_mistralServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_mistralServiceSettings) ApiKey(apikey string) *_mistralServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_mistralServiceSettings) MaxInputTokens(maxinputtokens int) *_mistralServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_mistralServiceSettings) Model(model string) *_mistralServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_mistralServiceSettings) RateLimit(ratelimit types.RateLimitSettingVariant) *_mistralServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_mistralServiceSettings) MistralServiceSettingsCaster() *types.MistralServiceSettings {
	_ = "STUB: not implemented"
	return nil
}
