package types

type GcsRepositorySettings struct {
	ApplicationName *string `json:"application_name,omitempty"`

	BasePath *string `json:"base_path,omitempty"`

	Bucket string `json:"bucket"`

	ChunkSize ByteSize `json:"chunk_size,omitempty"`

	Client *string `json:"client,omitempty"`

	Compress *bool `json:"compress,omitempty"`

	MaxRestoreBytesPerSec ByteSize `json:"max_restore_bytes_per_sec,omitempty"`

	MaxSnapshotBytesPerSec ByteSize `json:"max_snapshot_bytes_per_sec,omitempty"`

	Readonly *bool `json:"readonly,omitempty"`
}

func (s *GcsRepositorySettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewGcsRepositorySettings() *GcsRepositorySettings { _ = "STUB: not implemented"; return nil }

type GcsRepositorySettingsVariant interface {
	GcsRepositorySettingsCaster() *GcsRepositorySettings
}

func (s *GcsRepositorySettings) GcsRepositorySettingsCaster() *GcsRepositorySettings {
	_ = "STUB: not implemented"
	return nil
}
