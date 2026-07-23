package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _alibabaCloudServiceSettings struct {
	v *types.AlibabaCloudServiceSettings
}

func NewAlibabaCloudServiceSettings(apikey string, host string, serviceid string, workspace string) *_alibabaCloudServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_alibabaCloudServiceSettings) ApiKey(apikey string) *_alibabaCloudServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_alibabaCloudServiceSettings) Host(host string) *_alibabaCloudServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_alibabaCloudServiceSettings) RateLimit(ratelimit types.RateLimitSettingVariant) *_alibabaCloudServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_alibabaCloudServiceSettings) ServiceId(serviceid string) *_alibabaCloudServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_alibabaCloudServiceSettings) Workspace(workspace string) *_alibabaCloudServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_alibabaCloudServiceSettings) AlibabaCloudServiceSettingsCaster() *types.AlibabaCloudServiceSettings {
	_ = "STUB: not implemented"
	return nil
}
