package types

type ThinkingConfig struct {
	ThinkingBudget *int `json:"thinking_budget,omitempty"`
}

func (s *ThinkingConfig) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewThinkingConfig() *ThinkingConfig { _ = "STUB: not implemented"; return nil }

type ThinkingConfigVariant interface {
	ThinkingConfigCaster() *ThinkingConfig
}

func (s *ThinkingConfig) ThinkingConfigCaster() *ThinkingConfig {
	_ = "STUB: not implemented"
	return nil
}
