package types

type WatcherStatusActions map[string]ActionStatus

type WatcherStatusActionsVariant interface {
	WatcherStatusActionsCaster() *WatcherStatusActions
}
