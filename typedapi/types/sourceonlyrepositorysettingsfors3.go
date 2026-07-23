package types

type SourceOnlyRepositorySettingsForS3 struct {
	BasePath *string `json:"base_path,omitempty"`

	Bucket string `json:"bucket"`

	BufferSize ByteSize `json:"buffer_size,omitempty"`

	CannedAcl *string `json:"canned_acl,omitempty"`

	Client       *string `json:"client,omitempty"`
	DelegateType string  `json:"delegate_type,omitempty"`

	DeleteObjectsMaxSize *int `json:"delete_objects_max_size,omitempty"`

	GetRegisterRetryDelay Duration `json:"get_register_retry_delay,omitempty"`

	MaxMultipartParts *int `json:"max_multipart_parts,omitempty"`

	MaxMultipartUploadCleanupSize *int `json:"max_multipart_upload_cleanup_size,omitempty"`

	Readonly *bool `json:"readonly,omitempty"`

	ServerSideEncryption *bool `json:"server_side_encryption,omitempty"`

	StorageClass *string `json:"storage_class,omitempty"`

	ThrottledDeleteRetryDelayIncrement Duration `json:"throttled_delete_retry.delay_increment,omitempty"`

	ThrottledDeleteRetryMaximumDelay Duration `json:"throttled_delete_retry.maximum_delay,omitempty"`

	ThrottledDeleteRetryMaximumNumberOfRetries *int `json:"throttled_delete_retry.maximum_number_of_retries,omitempty"`
}

func (s *SourceOnlyRepositorySettingsForS3) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s SourceOnlyRepositorySettingsForS3) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewSourceOnlyRepositorySettingsForS3() *SourceOnlyRepositorySettingsForS3 {
	_ = "STUB: not implemented"
	return nil
}

type SourceOnlyRepositorySettingsForS3Variant interface {
	SourceOnlyRepositorySettingsForS3Caster() *SourceOnlyRepositorySettingsForS3
}

func (s *SourceOnlyRepositorySettingsForS3) SourceOnlyRepositorySettingsForS3Caster() *SourceOnlyRepositorySettingsForS3 {
	_ = "STUB: not implemented"
	return nil
}

func (s *SourceOnlyRepositorySettingsForS3) SourceOnlyRepositorySettingsCaster() *SourceOnlyRepositorySettings {
	_ = "STUB: not implemented"
	return nil
}
