package types

type Settings struct {
	AlignCheckpoints *bool `json:"align_checkpoints,omitempty"`

	DatesAsEpochMillis *bool `json:"dates_as_epoch_millis,omitempty"`

	DeduceMappings *bool `json:"deduce_mappings,omitempty"`

	DocsPerSecond *float32 `json:"docs_per_second,omitempty"`

	MaxPageSearchSize *int `json:"max_page_search_size,omitempty"`

	NumFailureRetries *int `json:"num_failure_retries,omitempty"`

	Unattended *bool `json:"unattended,omitempty"`

	UsePointInTime *bool `json:"use_point_in_time,omitempty"`
}

func (s *Settings) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSettings() *Settings { _ = "STUB: not implemented"; return nil }

type SettingsVariant interface {
	SettingsCaster() *Settings
}

func (s *Settings) SettingsCaster() *Settings { _ = "STUB: not implemented"; return nil }
