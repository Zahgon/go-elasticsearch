package types

type WatcherActions struct {
	Actions map[string]WatcherActionTotals `json:"actions"`
}

func NewWatcherActions() *WatcherActions { _ = "STUB: not implemented"; return nil }
