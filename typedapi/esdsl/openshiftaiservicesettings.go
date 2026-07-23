package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/openshiftaisimilaritytype"
)

type _openShiftAiServiceSettings struct {
	v *types.OpenShiftAiServiceSettings
}

func NewOpenShiftAiServiceSettings(apikey string, url string) *_openShiftAiServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_openShiftAiServiceSettings) ApiKey(apikey string) *_openShiftAiServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_openShiftAiServiceSettings) MaxInputTokens(maxinputtokens int) *_openShiftAiServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_openShiftAiServiceSettings) ModelId(modelid string) *_openShiftAiServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_openShiftAiServiceSettings) RateLimit(ratelimit types.RateLimitSettingVariant) *_openShiftAiServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_openShiftAiServiceSettings) Similarity(similarity openshiftaisimilaritytype.OpenShiftAiSimilarityType) *_openShiftAiServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_openShiftAiServiceSettings) Url(url string) *_openShiftAiServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_openShiftAiServiceSettings) OpenShiftAiServiceSettingsCaster() *types.OpenShiftAiServiceSettings {
	_ = "STUB: not implemented"
	return nil
}
