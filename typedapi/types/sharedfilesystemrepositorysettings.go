package types

type SharedFileSystemRepositorySettings struct {
	ChunkSize ByteSize `json:"chunk_size,omitempty"`

	Compress *bool `json:"compress,omitempty"`

	Location string `json:"location"`

	MaxNumberOfSnapshots *int `json:"max_number_of_snapshots,omitempty"`

	MaxRestoreBytesPerSec ByteSize `json:"max_restore_bytes_per_sec,omitempty"`

	MaxSnapshotBytesPerSec ByteSize `json:"max_snapshot_bytes_per_sec,omitempty"`

	Readonly *bool `json:"readonly,omitempty"`
}

func (s *SharedFileSystemRepositorySettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewSharedFileSystemRepositorySettings() *SharedFileSystemRepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

type SharedFileSystemRepositorySettingsVariant interface {
	SharedFileSystemRepositorySettingsCaster() *SharedFileSystemRepositorySettings
}

func (s *SharedFileSystemRepositorySettings) SharedFileSystemRepositorySettingsCaster() *SharedFileSystemRepositorySettings {
	_ = "STUB: not implemented"
	return nil
}
