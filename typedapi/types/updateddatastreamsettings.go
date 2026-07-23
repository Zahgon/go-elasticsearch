package types

type UpdatedDataStreamSettings struct {
	AppliedToDataStream bool `json:"applied_to_data_stream"`

	EffectiveSettings IndexSettings `json:"effective_settings"`

	Error *string `json:"error,omitempty"`

	IndexSettingsResults IndexSettingResults `json:"index_settings_results"`

	Name string `json:"name"`

	Settings IndexSettings `json:"settings"`
}

func (s *UpdatedDataStreamSettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewUpdatedDataStreamSettings() *UpdatedDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}
