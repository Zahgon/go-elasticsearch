package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _watchStatus struct {
	v *types.WatchStatus
}

func NewWatchStatus(state types.ActivationStateVariant) *_watchStatus {
	_ = "STUB: not implemented"
	return nil
}

func (s *_watchStatus) Actions(watcherstatusactions types.WatcherStatusActionsVariant) *_watchStatus {
	_ = "STUB: not implemented"
	return nil
}

func (s *_watchStatus) ExecutionState(executionstate string) *_watchStatus {
	_ = "STUB: not implemented"
	return nil
}

func (s *_watchStatus) LastChecked(datetime types.DateTimeVariant) *_watchStatus {
	_ = "STUB: not implemented"
	return nil
}

func (s *_watchStatus) LastMetCondition(datetime types.DateTimeVariant) *_watchStatus {
	_ = "STUB: not implemented"
	return nil
}

func (s *_watchStatus) State(state types.ActivationStateVariant) *_watchStatus {
	_ = "STUB: not implemented"
	return nil
}

func (s *_watchStatus) Version(versionnumber int64) *_watchStatus {
	_ = "STUB: not implemented"
	return nil
}

func (s *_watchStatus) WatchStatusCaster() *types.WatchStatus {
	_ = "STUB: not implemented"
	return nil
}
