package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _watsonxServiceSettings struct {
	v *types.WatsonxServiceSettings
}

func NewWatsonxServiceSettings(apikey string, apiversion string, modelid string, projectid string, url string) *_watsonxServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_watsonxServiceSettings) ApiKey(apikey string) *_watsonxServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_watsonxServiceSettings) ApiVersion(apiversion string) *_watsonxServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_watsonxServiceSettings) ModelId(modelid string) *_watsonxServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_watsonxServiceSettings) ProjectId(projectid string) *_watsonxServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_watsonxServiceSettings) RateLimit(ratelimit types.RateLimitSettingVariant) *_watsonxServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_watsonxServiceSettings) Url(url string) *_watsonxServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_watsonxServiceSettings) WatsonxServiceSettingsCaster() *types.WatsonxServiceSettings {
	_ = "STUB: not implemented"
	return nil
}
