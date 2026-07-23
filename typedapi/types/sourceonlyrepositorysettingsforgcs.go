package types

type SourceOnlyRepositorySettingsForGcs struct {
	ApplicationName *string `json:"application_name,omitempty"`

	BasePath *string `json:"base_path,omitempty"`

	Bucket string `json:"bucket"`

	Client       *string `json:"client,omitempty"`
	DelegateType string  `json:"delegate_type,omitempty"`

	Readonly *bool `json:"readonly,omitempty"`
}

func (s *SourceOnlyRepositorySettingsForGcs) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s SourceOnlyRepositorySettingsForGcs) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewSourceOnlyRepositorySettingsForGcs() *SourceOnlyRepositorySettingsForGcs {
	_ = "STUB: not implemented"
	return nil
}

type SourceOnlyRepositorySettingsForGcsVariant interface {
	SourceOnlyRepositorySettingsForGcsCaster() *SourceOnlyRepositorySettingsForGcs
}

func (s *SourceOnlyRepositorySettingsForGcs) SourceOnlyRepositorySettingsForGcsCaster() *SourceOnlyRepositorySettingsForGcs {
	_ = "STUB: not implemented"
	return nil
}

func (s *SourceOnlyRepositorySettingsForGcs) SourceOnlyRepositorySettingsCaster() *SourceOnlyRepositorySettings {
	_ = "STUB: not implemented"
	return nil
}
