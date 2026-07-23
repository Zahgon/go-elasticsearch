package types

type SettingsAnalyze struct {
	MaxTokenCount Stringifiedinteger `json:"max_token_count,omitempty"`
}

func (s *SettingsAnalyze) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSettingsAnalyze() *SettingsAnalyze { _ = "STUB: not implemented"; return nil }

type SettingsAnalyzeVariant interface {
	SettingsAnalyzeCaster() *SettingsAnalyze
}

func (s *SettingsAnalyze) SettingsAnalyzeCaster() *SettingsAnalyze {
	_ = "STUB: not implemented"
	return nil
}
