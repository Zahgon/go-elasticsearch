package types

type RateLimitSetting struct {
	RequestsPerMinute *int `json:"requests_per_minute,omitempty"`
}

func (s *RateLimitSetting) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRateLimitSetting() *RateLimitSetting { _ = "STUB: not implemented"; return nil }

type RateLimitSettingVariant interface {
	RateLimitSettingCaster() *RateLimitSetting
}

func (s *RateLimitSetting) RateLimitSettingCaster() *RateLimitSetting {
	_ = "STUB: not implemented"
	return nil
}
