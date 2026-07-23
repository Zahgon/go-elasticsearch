package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _syncContainer struct {
	v *types.SyncContainer
}

func NewSyncContainer() *_syncContainer { _ = "STUB: not implemented"; return nil }

func (s *_syncContainer) Time(time types.TimeSyncVariant) *_syncContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_syncContainer) SyncContainerCaster() *types.SyncContainer {
	_ = "STUB: not implemented"
	return nil
}
