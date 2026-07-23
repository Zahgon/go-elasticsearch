package types

type GoogleVertexAITaskSettings struct {
	AutoTruncate *bool `json:"auto_truncate,omitempty"`

	MaxTokens *int `json:"max_tokens,omitempty"`

	ThinkingConfig *ThinkingConfig `json:"thinking_config,omitempty"`

	TopN *int `json:"top_n,omitempty"`
}

func (s *GoogleVertexAITaskSettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewGoogleVertexAITaskSettings() *GoogleVertexAITaskSettings {
	_ = "STUB: not implemented"
	return nil
}

type GoogleVertexAITaskSettingsVariant interface {
	GoogleVertexAITaskSettingsCaster() *GoogleVertexAITaskSettings
}

func (s *GoogleVertexAITaskSettings) GoogleVertexAITaskSettingsCaster() *GoogleVertexAITaskSettings {
	_ = "STUB: not implemented"
	return nil
}
