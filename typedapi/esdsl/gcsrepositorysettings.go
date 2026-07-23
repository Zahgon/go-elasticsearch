package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _gcsRepositorySettings struct {
	v *types.GcsRepositorySettings
}

func NewGcsRepositorySettings(bucket string) *_gcsRepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_gcsRepositorySettings) ApplicationName(applicationname string) *_gcsRepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_gcsRepositorySettings) BasePath(basepath string) *_gcsRepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_gcsRepositorySettings) Bucket(bucket string) *_gcsRepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_gcsRepositorySettings) Client(client string) *_gcsRepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_gcsRepositorySettings) Readonly(readonly bool) *_gcsRepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_gcsRepositorySettings) ChunkSize(bytesize types.ByteSizeVariant) *_gcsRepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_gcsRepositorySettings) Compress(compress bool) *_gcsRepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_gcsRepositorySettings) MaxRestoreBytesPerSec(bytesize types.ByteSizeVariant) *_gcsRepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_gcsRepositorySettings) MaxSnapshotBytesPerSec(bytesize types.ByteSizeVariant) *_gcsRepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_gcsRepositorySettings) GcsRepositorySettingsCaster() *types.GcsRepositorySettings {
	_ = "STUB: not implemented"
	return nil
}
