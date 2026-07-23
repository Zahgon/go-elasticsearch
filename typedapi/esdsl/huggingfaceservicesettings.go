package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _huggingFaceServiceSettings struct {
	v *types.HuggingFaceServiceSettings
}

func NewHuggingFaceServiceSettings(apikey string, url string) *_huggingFaceServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_huggingFaceServiceSettings) ApiKey(apikey string) *_huggingFaceServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_huggingFaceServiceSettings) ModelId(modelid string) *_huggingFaceServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_huggingFaceServiceSettings) RateLimit(ratelimit types.RateLimitSettingVariant) *_huggingFaceServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_huggingFaceServiceSettings) Url(url string) *_huggingFaceServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_huggingFaceServiceSettings) HuggingFaceServiceSettingsCaster() *types.HuggingFaceServiceSettings {
	_ = "STUB: not implemented"
	return nil
}
