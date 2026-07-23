package types

type ReadOnlyUrlRepositorySettings struct {
	ChunkSize ByteSize `json:"chunk_size,omitempty"`

	Compress *bool `json:"compress,omitempty"`

	HttpMaxRetries *int `json:"http_max_retries,omitempty"`

	HttpSocketTimeout Duration `json:"http_socket_timeout,omitempty"`

	MaxNumberOfSnapshots *int `json:"max_number_of_snapshots,omitempty"`

	MaxRestoreBytesPerSec ByteSize `json:"max_restore_bytes_per_sec,omitempty"`

	MaxSnapshotBytesPerSec ByteSize `json:"max_snapshot_bytes_per_sec,omitempty"`

	Url string `json:"url"`
}

func (s *ReadOnlyUrlRepositorySettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewReadOnlyUrlRepositorySettings() *ReadOnlyUrlRepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

type ReadOnlyUrlRepositorySettingsVariant interface {
	ReadOnlyUrlRepositorySettingsCaster() *ReadOnlyUrlRepositorySettings
}

func (s *ReadOnlyUrlRepositorySettings) ReadOnlyUrlRepositorySettingsCaster() *ReadOnlyUrlRepositorySettings {
	_ = "STUB: not implemented"
	return nil
}
