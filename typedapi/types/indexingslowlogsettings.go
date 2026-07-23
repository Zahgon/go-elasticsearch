package types

type IndexingSlowlogSettings struct {
	Level     *string                   `json:"level,omitempty"`
	Reformat  *bool                     `json:"reformat,omitempty"`
	Source    *int                      `json:"source,omitempty"`
	Threshold *IndexingSlowlogTresholds `json:"threshold,omitempty"`
}

func (s *IndexingSlowlogSettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewIndexingSlowlogSettings() *IndexingSlowlogSettings { _ = "STUB: not implemented"; return nil }

type IndexingSlowlogSettingsVariant interface {
	IndexingSlowlogSettingsCaster() *IndexingSlowlogSettings
}

func (s *IndexingSlowlogSettings) IndexingSlowlogSettingsCaster() *IndexingSlowlogSettings {
	_ = "STUB: not implemented"
	return nil
}
