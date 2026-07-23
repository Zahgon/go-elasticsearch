package types

type CatComponentTemplate struct {
	AliasCount    string  `json:"alias_count"`
	IncludedIn    string  `json:"included_in"`
	MappingCount  string  `json:"mapping_count"`
	MetadataCount string  `json:"metadata_count"`
	Name          string  `json:"name"`
	SettingsCount string  `json:"settings_count"`
	Version       *string `json:"version,omitempty"`
}

func (s *CatComponentTemplate) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewCatComponentTemplate() *CatComponentTemplate { _ = "STUB: not implemented"; return nil }
