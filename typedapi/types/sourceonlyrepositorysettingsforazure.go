package types

type SourceOnlyRepositorySettingsForAzure struct {
	BasePath *string `json:"base_path,omitempty"`

	Client *string `json:"client,omitempty"`

	Container    *string `json:"container,omitempty"`
	DelegateType string  `json:"delegate_type,omitempty"`

	DeleteObjectsMaxSize *int `json:"delete_objects_max_size,omitempty"`

	LocationMode *string `json:"location_mode,omitempty"`

	MaxConcurrentBatchDeletes *int `json:"max_concurrent_batch_deletes,omitempty"`

	Readonly *bool `json:"readonly,omitempty"`
}

func (s *SourceOnlyRepositorySettingsForAzure) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s SourceOnlyRepositorySettingsForAzure) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewSourceOnlyRepositorySettingsForAzure() *SourceOnlyRepositorySettingsForAzure {
	_ = "STUB: not implemented"
	return nil
}

type SourceOnlyRepositorySettingsForAzureVariant interface {
	SourceOnlyRepositorySettingsForAzureCaster() *SourceOnlyRepositorySettingsForAzure
}

func (s *SourceOnlyRepositorySettingsForAzure) SourceOnlyRepositorySettingsForAzureCaster() *SourceOnlyRepositorySettingsForAzure {
	_ = "STUB: not implemented"
	return nil
}

func (s *SourceOnlyRepositorySettingsForAzure) SourceOnlyRepositorySettingsCaster() *SourceOnlyRepositorySettings {
	_ = "STUB: not implemented"
	return nil
}
