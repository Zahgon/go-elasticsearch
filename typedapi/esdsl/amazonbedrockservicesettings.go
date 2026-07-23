package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _amazonBedrockServiceSettings struct {
	v *types.AmazonBedrockServiceSettings
}

func NewAmazonBedrockServiceSettings(accesskey string, model string, region string, secretkey string) *_amazonBedrockServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_amazonBedrockServiceSettings) AccessKey(accesskey string) *_amazonBedrockServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_amazonBedrockServiceSettings) Model(model string) *_amazonBedrockServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_amazonBedrockServiceSettings) Provider(provider string) *_amazonBedrockServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_amazonBedrockServiceSettings) RateLimit(ratelimit types.RateLimitSettingVariant) *_amazonBedrockServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_amazonBedrockServiceSettings) Region(region string) *_amazonBedrockServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_amazonBedrockServiceSettings) SecretKey(secretkey string) *_amazonBedrockServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_amazonBedrockServiceSettings) AmazonBedrockServiceSettingsCaster() *types.AmazonBedrockServiceSettings {
	_ = "STUB: not implemented"
	return nil
}
