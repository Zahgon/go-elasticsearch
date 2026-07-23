package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/storagetype"
)

type _storage struct {
	v *types.Storage
}

func NewStorage(type_ storagetype.StorageType) *_storage { _ = "STUB: not implemented"; return nil }

func (s *_storage) AllowMmap(allowmmap bool) *_storage { _ = "STUB: not implemented"; return nil }

func (s *_storage) StatsRefreshInterval(duration types.DurationVariant) *_storage {
	_ = "STUB: not implemented"
	return nil
}

func (s *_storage) Type(type_ storagetype.StorageType) *_storage {
	_ = "STUB: not implemented"
	return nil
}

func (s *_storage) StorageCaster() *types.Storage { _ = "STUB: not implemented"; return nil }
