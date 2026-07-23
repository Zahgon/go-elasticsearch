package types

type AlibabaCloudTaskSettings struct {
	InputType *string `json:"input_type,omitempty"`

	ReturnToken *bool `json:"return_token,omitempty"`
}

func (s *AlibabaCloudTaskSettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewAlibabaCloudTaskSettings() *AlibabaCloudTaskSettings { _ = "STUB: not implemented"; return nil }

type AlibabaCloudTaskSettingsVariant interface {
	AlibabaCloudTaskSettingsCaster() *AlibabaCloudTaskSettings
}

func (s *AlibabaCloudTaskSettings) AlibabaCloudTaskSettingsCaster() *AlibabaCloudTaskSettings {
	_ = "STUB: not implemented"
	return nil
}
