package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _readOnlyUrlRepositorySettings struct {
	v *types.ReadOnlyUrlRepositorySettings
}

func NewReadOnlyUrlRepositorySettings(url string) *_readOnlyUrlRepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_readOnlyUrlRepositorySettings) HttpMaxRetries(httpmaxretries int) *_readOnlyUrlRepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_readOnlyUrlRepositorySettings) HttpSocketTimeout(duration types.DurationVariant) *_readOnlyUrlRepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_readOnlyUrlRepositorySettings) MaxNumberOfSnapshots(maxnumberofsnapshots int) *_readOnlyUrlRepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_readOnlyUrlRepositorySettings) Url(url string) *_readOnlyUrlRepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_readOnlyUrlRepositorySettings) ChunkSize(bytesize types.ByteSizeVariant) *_readOnlyUrlRepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_readOnlyUrlRepositorySettings) Compress(compress bool) *_readOnlyUrlRepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_readOnlyUrlRepositorySettings) MaxRestoreBytesPerSec(bytesize types.ByteSizeVariant) *_readOnlyUrlRepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_readOnlyUrlRepositorySettings) MaxSnapshotBytesPerSec(bytesize types.ByteSizeVariant) *_readOnlyUrlRepositorySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_readOnlyUrlRepositorySettings) ReadOnlyUrlRepositorySettingsCaster() *types.ReadOnlyUrlRepositorySettings {
	_ = "STUB: not implemented"
	return nil
}
