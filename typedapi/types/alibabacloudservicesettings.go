package types

type AlibabaCloudServiceSettings struct {
	ApiKey string `json:"api_key"`

	Host string `json:"host"`

	RateLimit *RateLimitSetting `json:"rate_limit,omitempty"`

	ServiceId string `json:"service_id"`

	Workspace string `json:"workspace"`
}

func (s *AlibabaCloudServiceSettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewAlibabaCloudServiceSettings() *AlibabaCloudServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

type AlibabaCloudServiceSettingsVariant interface {
	AlibabaCloudServiceSettingsCaster() *AlibabaCloudServiceSettings
}

func (s *AlibabaCloudServiceSettings) AlibabaCloudServiceSettingsCaster() *AlibabaCloudServiceSettings {
	_ = "STUB: not implemented"
	return nil
}
