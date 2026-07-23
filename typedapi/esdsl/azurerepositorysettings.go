package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _azureRepositorySettings struct {
	v *types.AzureRepositorySettings
}

func NewAzureRepositorySettings() *_azureRepositorySettings { _ = "STUB: not implemented"; return nil }

func (s *_azureRepositorySettings) BasePath(basepath string) *_azureRepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_azureRepositorySettings) Client(client string) *_azureRepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_azureRepositorySettings) Container(container string) *_azureRepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_azureRepositorySettings) DeleteObjectsMaxSize(deleteobjectsmaxsize int) *_azureRepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_azureRepositorySettings) LocationMode(locationmode string) *_azureRepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_azureRepositorySettings) MaxConcurrentBatchDeletes(maxconcurrentbatchdeletes int) *_azureRepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_azureRepositorySettings) Readonly(readonly bool) *_azureRepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_azureRepositorySettings) ChunkSize(bytesize types.ByteSizeVariant) *_azureRepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_azureRepositorySettings) Compress(compress bool) *_azureRepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_azureRepositorySettings) MaxRestoreBytesPerSec(bytesize types.ByteSizeVariant) *_azureRepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_azureRepositorySettings) MaxSnapshotBytesPerSec(bytesize types.ByteSizeVariant) *_azureRepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_azureRepositorySettings) AzureRepositorySettingsCaster() *types.AzureRepositorySettings {
	_ = "STUB: not implemented"
	return nil
}
