package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _rateLimitSetting struct {
	v *types.RateLimitSetting
}

func NewRateLimitSetting() *_rateLimitSetting { _ = "STUB: not implemented"; return nil }

func (s *_rateLimitSetting) RequestsPerMinute(requestsperminute int) *_rateLimitSetting {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rateLimitSetting) RateLimitSettingCaster() *types.RateLimitSetting {
	_ = "STUB: not implemented"
	return nil
}
