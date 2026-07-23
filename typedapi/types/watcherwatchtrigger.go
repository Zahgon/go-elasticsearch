package types

type WatcherWatchTrigger struct {
	All_     Counter                      `json:"_all"`
	Schedule *WatcherWatchTriggerSchedule `json:"schedule,omitempty"`
}

func NewWatcherWatchTrigger() *WatcherWatchTrigger { _ = "STUB: not implemented"; return nil }
