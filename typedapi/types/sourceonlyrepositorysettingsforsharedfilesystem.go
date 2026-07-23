package types

type SourceOnlyRepositorySettingsForSharedFileSystem struct {
	DelegateType string `json:"delegate_type,omitempty"`

	Location string `json:"location"`

	MaxNumberOfSnapshots *int `json:"max_number_of_snapshots,omitempty"`

	Readonly *bool `json:"readonly,omitempty"`
}

func (s *SourceOnlyRepositorySettingsForSharedFileSystem) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s SourceOnlyRepositorySettingsForSharedFileSystem) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewSourceOnlyRepositorySettingsForSharedFileSystem() *SourceOnlyRepositorySettingsForSharedFileSystem {
	_ = "STUB: not implemented"
	return nil
}

type SourceOnlyRepositorySettingsForSharedFileSystemVariant interface {
	SourceOnlyRepositorySettingsForSharedFileSystemCaster() *SourceOnlyRepositorySettingsForSharedFileSystem
}

func (s *SourceOnlyRepositorySettingsForSharedFileSystem) SourceOnlyRepositorySettingsForSharedFileSystemCaster() *SourceOnlyRepositorySettingsForSharedFileSystem {
	_ = "STUB: not implemented"
	return nil
}

func (s *SourceOnlyRepositorySettingsForSharedFileSystem) SourceOnlyRepositorySettingsCaster() *SourceOnlyRepositorySettings {
	_ = "STUB: not implemented"
	return nil
}
