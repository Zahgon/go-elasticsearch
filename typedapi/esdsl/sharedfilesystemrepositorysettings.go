package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _sharedFileSystemRepositorySettings struct {
	v *types.SharedFileSystemRepositorySettings
}

func NewSharedFileSystemRepositorySettings(location string) *_sharedFileSystemRepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sharedFileSystemRepositorySettings) Location(location string) *_sharedFileSystemRepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sharedFileSystemRepositorySettings) MaxNumberOfSnapshots(maxnumberofsnapshots int) *_sharedFileSystemRepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sharedFileSystemRepositorySettings) Readonly(readonly bool) *_sharedFileSystemRepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sharedFileSystemRepositorySettings) ChunkSize(bytesize types.ByteSizeVariant) *_sharedFileSystemRepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sharedFileSystemRepositorySettings) Compress(compress bool) *_sharedFileSystemRepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sharedFileSystemRepositorySettings) MaxRestoreBytesPerSec(bytesize types.ByteSizeVariant) *_sharedFileSystemRepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sharedFileSystemRepositorySettings) MaxSnapshotBytesPerSec(bytesize types.ByteSizeVariant) *_sharedFileSystemRepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sharedFileSystemRepositorySettings) SharedFileSystemRepositorySettingsCaster() *types.SharedFileSystemRepositorySettings {
	_ = "STUB: not implemented"
	return nil
}
