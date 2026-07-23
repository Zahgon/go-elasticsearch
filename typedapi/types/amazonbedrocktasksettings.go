package types

type AmazonBedrockTaskSettings struct {
	MaxNewTokens *int `json:"max_new_tokens,omitempty"`

	Temperature *float32 `json:"temperature,omitempty"`

	TopK *float32 `json:"top_k,omitempty"`

	TopP *float32 `json:"top_p,omitempty"`
}

func (s *AmazonBedrockTaskSettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewAmazonBedrockTaskSettings() *AmazonBedrockTaskSettings {
	_ = "STUB: not implemented"
	return nil
}

type AmazonBedrockTaskSettingsVariant interface {
	AmazonBedrockTaskSettingsCaster() *AmazonBedrockTaskSettings
}

func (s *AmazonBedrockTaskSettings) AmazonBedrockTaskSettingsCaster() *AmazonBedrockTaskSettings {
	_ = "STUB: not implemented"
	return nil
}
