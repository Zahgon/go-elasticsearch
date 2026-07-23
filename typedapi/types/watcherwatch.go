package types

type WatcherWatch struct {
	Action    map[string]Counter  `json:"action,omitempty"`
	Condition map[string]Counter  `json:"condition,omitempty"`
	Input     map[string]Counter  `json:"input"`
	Trigger   WatcherWatchTrigger `json:"trigger"`
}

func NewWatcherWatch() *WatcherWatch { _ = "STUB: not implemented"; return nil }
