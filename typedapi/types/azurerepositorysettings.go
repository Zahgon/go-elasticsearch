package types

type AzureRepositorySettings struct {
	BasePath *string `json:"base_path,omitempty"`

	ChunkSize ByteSize `json:"chunk_size,omitempty"`

	Client *string `json:"client,omitempty"`

	Compress *bool `json:"compress,omitempty"`

	Container *string `json:"container,omitempty"`

	DeleteObjectsMaxSize *int `json:"delete_objects_max_size,omitempty"`

	LocationMode *string `json:"location_mode,omitempty"`

	MaxConcurrentBatchDeletes *int `json:"max_concurrent_batch_deletes,omitempty"`

	MaxRestoreBytesPerSec ByteSize `json:"max_restore_bytes_per_sec,omitempty"`

	MaxSnapshotBytesPerSec ByteSize `json:"max_snapshot_bytes_per_sec,omitempty"`

	Readonly *bool `json:"readonly,omitempty"`
}

func (s *AzureRepositorySettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewAzureRepositorySettings() *AzureRepositorySettings { _ = "STUB: not implemented"; return nil }

type AzureRepositorySettingsVariant interface {
	AzureRepositorySettingsCaster() *AzureRepositorySettings
}

func (s *AzureRepositorySettings) AzureRepositorySettingsCaster() *AzureRepositorySettings {
	_ = "STUB: not implemented"
	return nil
}
