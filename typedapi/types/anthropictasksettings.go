package types

type AnthropicTaskSettings struct {
	MaxTokens int `json:"max_tokens"`

	Temperature *float32 `json:"temperature,omitempty"`

	TopK *int `json:"top_k,omitempty"`

	TopP *float32 `json:"top_p,omitempty"`
}

func (s *AnthropicTaskSettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewAnthropicTaskSettings() *AnthropicTaskSettings { _ = "STUB: not implemented"; return nil }

type AnthropicTaskSettingsVariant interface {
	AnthropicTaskSettingsCaster() *AnthropicTaskSettings
}

func (s *AnthropicTaskSettings) AnthropicTaskSettingsCaster() *AnthropicTaskSettings {
	_ = "STUB: not implemented"
	return nil
}
