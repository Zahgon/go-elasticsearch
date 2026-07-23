package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/watcherstate"
)

type WatcherNodeStats struct {
	CurrentWatches      []WatchRecordStats  `json:"current_watches,omitempty"`
	ExecutionThreadPool ExecutionThreadPool `json:"execution_thread_pool"`
	NodeId              string              `json:"node_id"`

	QueuedWatches []WatchRecordQueuedStats `json:"queued_watches,omitempty"`

	WatchCount int64 `json:"watch_count"`

	WatcherState watcherstate.WatcherState `json:"watcher_state"`
}

func (s *WatcherNodeStats) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewWatcherNodeStats() *WatcherNodeStats { _ = "STUB: not implemented"; return nil }
