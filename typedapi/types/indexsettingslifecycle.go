package types

type IndexSettingsLifecycle struct {
	IndexingComplete Stringifiedboolean `json:"indexing_complete,omitempty"`

	Name *string `json:"name,omitempty"`

	OriginationDate *int64 `json:"origination_date,omitempty"`

	ParseOriginationDate *bool `json:"parse_origination_date,omitempty"`

	PreferIlm *string `json:"prefer_ilm,omitempty"`

	RolloverAlias *string                     `json:"rollover_alias,omitempty"`
	Step          *IndexSettingsLifecycleStep `json:"step,omitempty"`
}

func (s *IndexSettingsLifecycle) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewIndexSettingsLifecycle() *IndexSettingsLifecycle { _ = "STUB: not implemented"; return nil }

type IndexSettingsLifecycleVariant interface {
	IndexSettingsLifecycleCaster() *IndexSettingsLifecycle
}

func (s *IndexSettingsLifecycle) IndexSettingsLifecycleCaster() *IndexSettingsLifecycle {
	_ = "STUB: not implemented"
	return nil
}
