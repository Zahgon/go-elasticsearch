package watchermetric

type WatcherMetric struct {
	Name string
}

var (
	All = WatcherMetric{"_all"}

	Queuedwatches = WatcherMetric{"queued_watches"}

	Currentwatches = WatcherMetric{"current_watches"}

	Pendingwatches = WatcherMetric{"pending_watches"}
)

func (w WatcherMetric) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *WatcherMetric) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (w WatcherMetric) String() string { _ = "STUB: not implemented"; return "" }
