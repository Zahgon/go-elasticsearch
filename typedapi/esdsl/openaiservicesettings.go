package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/openaisimilaritytype"
)

type _openAIServiceSettings struct {
	v *types.OpenAIServiceSettings
}

func NewOpenAIServiceSettings(modelid string) *_openAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_openAIServiceSettings) ApiKey(apikey string) *_openAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_openAIServiceSettings) ClientId(clientid string) *_openAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_openAIServiceSettings) ClientSecret(clientsecret string) *_openAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_openAIServiceSettings) Dimensions(dimensions int) *_openAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_openAIServiceSettings) ModelId(modelid string) *_openAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_openAIServiceSettings) OrganizationId(organizationid string) *_openAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_openAIServiceSettings) RateLimit(ratelimit types.RateLimitSettingVariant) *_openAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_openAIServiceSettings) Scopes(scopes ...string) *_openAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_openAIServiceSettings) Similarity(similarity openaisimilaritytype.OpenAISimilarityType) *_openAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_openAIServiceSettings) TokenUrl(tokenurl string) *_openAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_openAIServiceSettings) Url(url string) *_openAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_openAIServiceSettings) OpenAIServiceSettingsCaster() *types.OpenAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}
