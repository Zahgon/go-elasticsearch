package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _azureOpenAIServiceSettings struct {
	v *types.AzureOpenAIServiceSettings
}

func NewAzureOpenAIServiceSettings(apiversion string, deploymentid string, resourcename string) *_azureOpenAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_azureOpenAIServiceSettings) ApiKey(apikey string) *_azureOpenAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_azureOpenAIServiceSettings) ApiVersion(apiversion string) *_azureOpenAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_azureOpenAIServiceSettings) ClientId(clientid string) *_azureOpenAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_azureOpenAIServiceSettings) ClientSecret(clientsecret string) *_azureOpenAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_azureOpenAIServiceSettings) DeploymentId(deploymentid string) *_azureOpenAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_azureOpenAIServiceSettings) EntraId(entraid string) *_azureOpenAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_azureOpenAIServiceSettings) RateLimit(ratelimit types.RateLimitSettingVariant) *_azureOpenAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_azureOpenAIServiceSettings) ResourceName(resourcename string) *_azureOpenAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_azureOpenAIServiceSettings) Scopes(scopes ...string) *_azureOpenAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_azureOpenAIServiceSettings) TenantId(tenantid string) *_azureOpenAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_azureOpenAIServiceSettings) AzureOpenAIServiceSettingsCaster() *types.AzureOpenAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}
