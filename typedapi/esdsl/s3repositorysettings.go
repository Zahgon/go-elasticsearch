package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _s3RepositorySettings struct {
	v *types.S3RepositorySettings
}

func NewS3RepositorySettings(bucket string) *_s3RepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_s3RepositorySettings) BasePath(basepath string) *_s3RepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_s3RepositorySettings) Bucket(bucket string) *_s3RepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_s3RepositorySettings) BufferSize(bytesize types.ByteSizeVariant) *_s3RepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_s3RepositorySettings) CannedAcl(cannedacl string) *_s3RepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_s3RepositorySettings) Client(client string) *_s3RepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_s3RepositorySettings) DeleteObjectsMaxSize(deleteobjectsmaxsize int) *_s3RepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_s3RepositorySettings) GetRegisterRetryDelay(duration types.DurationVariant) *_s3RepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_s3RepositorySettings) MaxMultipartParts(maxmultipartparts int) *_s3RepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_s3RepositorySettings) MaxMultipartUploadCleanupSize(maxmultipartuploadcleanupsize int) *_s3RepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_s3RepositorySettings) Readonly(readonly bool) *_s3RepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_s3RepositorySettings) ServerSideEncryption(serversideencryption bool) *_s3RepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_s3RepositorySettings) StorageClass(storageclass string) *_s3RepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_s3RepositorySettings) ThrottledDeleteRetryDelayIncrement(duration types.DurationVariant) *_s3RepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_s3RepositorySettings) ThrottledDeleteRetryMaximumDelay(duration types.DurationVariant) *_s3RepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_s3RepositorySettings) ThrottledDeleteRetryMaximumNumberOfRetries(throttleddeleteretrymaximumnumberofretries int) *_s3RepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_s3RepositorySettings) ChunkSize(bytesize types.ByteSizeVariant) *_s3RepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_s3RepositorySettings) Compress(compress bool) *_s3RepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_s3RepositorySettings) MaxRestoreBytesPerSec(bytesize types.ByteSizeVariant) *_s3RepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_s3RepositorySettings) MaxSnapshotBytesPerSec(bytesize types.ByteSizeVariant) *_s3RepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_s3RepositorySettings) S3RepositorySettingsCaster() *types.S3RepositorySettings {
	_ = "STUB: not implemented"
	return nil
}
