package types

type HuggingFaceTaskSettings struct {
	ReturnDocuments *bool `json:"return_documents,omitempty"`

	TopN *int `json:"top_n,omitempty"`
}

func (s *HuggingFaceTaskSettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewHuggingFaceTaskSettings() *HuggingFaceTaskSettings { _ = "STUB: not implemented"; return nil }

type HuggingFaceTaskSettingsVariant interface {
	HuggingFaceTaskSettingsCaster() *HuggingFaceTaskSettings
}

func (s *HuggingFaceTaskSettings) HuggingFaceTaskSettingsCaster() *HuggingFaceTaskSettings {
	_ = "STUB: not implemented"
	return nil
}
