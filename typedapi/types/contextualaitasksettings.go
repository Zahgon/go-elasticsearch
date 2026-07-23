package types

type ContextualAITaskSettings struct {
	Instruction *string `json:"instruction,omitempty"`

	TopK *int `json:"top_k,omitempty"`
}

func (s *ContextualAITaskSettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewContextualAITaskSettings() *ContextualAITaskSettings { _ = "STUB: not implemented"; return nil }

type ContextualAITaskSettingsVariant interface {
	ContextualAITaskSettingsCaster() *ContextualAITaskSettings
}

func (s *ContextualAITaskSettings) ContextualAITaskSettingsCaster() *ContextualAITaskSettings {
	_ = "STUB: not implemented"
	return nil
}
