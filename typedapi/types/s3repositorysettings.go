package types

type S3RepositorySettings struct {
	BasePath *string `json:"base_path,omitempty"`

	Bucket string `json:"bucket"`

	BufferSize ByteSize `json:"buffer_size,omitempty"`

	CannedAcl *string `json:"canned_acl,omitempty"`

	ChunkSize ByteSize `json:"chunk_size,omitempty"`

	Client *string `json:"client,omitempty"`

	Compress *bool `json:"compress,omitempty"`

	DeleteObjectsMaxSize *int `json:"delete_objects_max_size,omitempty"`

	GetRegisterRetryDelay Duration `json:"get_register_retry_delay,omitempty"`

	MaxMultipartParts *int `json:"max_multipart_parts,omitempty"`

	MaxMultipartUploadCleanupSize *int `json:"max_multipart_upload_cleanup_size,omitempty"`

	MaxRestoreBytesPerSec ByteSize `json:"max_restore_bytes_per_sec,omitempty"`

	MaxSnapshotBytesPerSec ByteSize `json:"max_snapshot_bytes_per_sec,omitempty"`

	Readonly *bool `json:"readonly,omitempty"`

	ServerSideEncryption *bool `json:"server_side_encryption,omitempty"`

	StorageClass *string `json:"storage_class,omitempty"`

	ThrottledDeleteRetryDelayIncrement Duration `json:"throttled_delete_retry.delay_increment,omitempty"`

	ThrottledDeleteRetryMaximumDelay Duration `json:"throttled_delete_retry.maximum_delay,omitempty"`

	ThrottledDeleteRetryMaximumNumberOfRetries *int `json:"throttled_delete_retry.maximum_number_of_retries,omitempty"`
}

func (s *S3RepositorySettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewS3RepositorySettings() *S3RepositorySettings { _ = "STUB: not implemented"; return nil }

type S3RepositorySettingsVariant interface {
	S3RepositorySettingsCaster() *S3RepositorySettings
}

func (s *S3RepositorySettings) S3RepositorySettingsCaster() *S3RepositorySettings {
	_ = "STUB: not implemented"
	return nil
}
