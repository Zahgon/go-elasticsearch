package types

type VoyageAITaskSettings struct {
	InputType *string `json:"input_type,omitempty"`

	ReturnDocuments *bool `json:"return_documents,omitempty"`

	TopK *int `json:"top_k,omitempty"`

	Truncation *bool `json:"truncation,omitempty"`
}

func (s *VoyageAITaskSettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewVoyageAITaskSettings() *VoyageAITaskSettings { _ = "STUB: not implemented"; return nil }

type VoyageAITaskSettingsVariant interface {
	VoyageAITaskSettingsCaster() *VoyageAITaskSettings
}

func (s *VoyageAITaskSettings) VoyageAITaskSettingsCaster() *VoyageAITaskSettings {
	_ = "STUB: not implemented"
	return nil
}
