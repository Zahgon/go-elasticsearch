package types

type OpenShiftAiTaskSettings struct {
	ReturnDocuments *bool `json:"return_documents,omitempty"`

	TopN *int `json:"top_n,omitempty"`
}

func (s *OpenShiftAiTaskSettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewOpenShiftAiTaskSettings() *OpenShiftAiTaskSettings { _ = "STUB: not implemented"; return nil }

type OpenShiftAiTaskSettingsVariant interface {
	OpenShiftAiTaskSettingsCaster() *OpenShiftAiTaskSettings
}

func (s *OpenShiftAiTaskSettings) OpenShiftAiTaskSettingsCaster() *OpenShiftAiTaskSettings {
	_ = "STUB: not implemented"
	return nil
}
