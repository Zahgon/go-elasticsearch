package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _watcherStatusActions struct {
	v types.WatcherStatusActions
}

func NewWatcherStatusActions(watcherstatusactions map[string]types.ActionStatus) *_watcherStatusActions {
	_ = "STUB: not implemented"
	return nil
}

func (u *_watcherStatusActions) WatcherStatusActionsCaster() *types.WatcherStatusActions {
	_ = "STUB: not implemented"
	return nil
}
