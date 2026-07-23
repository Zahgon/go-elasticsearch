package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _azureAiStudioServiceSettings struct {
	v *types.AzureAiStudioServiceSettings
}

func NewAzureAiStudioServiceSettings(apikey string, endpointtype string, provider string, target string) *_azureAiStudioServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_azureAiStudioServiceSettings) ApiKey(apikey string) *_azureAiStudioServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_azureAiStudioServiceSettings) EndpointType(endpointtype string) *_azureAiStudioServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_azureAiStudioServiceSettings) Provider(provider string) *_azureAiStudioServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_azureAiStudioServiceSettings) RateLimit(ratelimit types.RateLimitSettingVariant) *_azureAiStudioServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_azureAiStudioServiceSettings) Target(target string) *_azureAiStudioServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_azureAiStudioServiceSettings) AzureAiStudioServiceSettingsCaster() *types.AzureAiStudioServiceSettings {
	_ = "STUB: not implemented"
	return nil
}
