package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _timeSync struct {
	v *types.TimeSync
}

func NewTimeSync() *_timeSync { _ = "STUB: not implemented"; return nil }

func (s *_timeSync) Delay(duration types.DurationVariant) *_timeSync {
	_ = "STUB: not implemented"
	return nil
}

func (s *_timeSync) Field(field string) *_timeSync { _ = "STUB: not implemented"; return nil }

func (s *_timeSync) SyncContainerCaster() *types.SyncContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_timeSync) TimeSyncCaster() *types.TimeSync { _ = "STUB: not implemented"; return nil }
