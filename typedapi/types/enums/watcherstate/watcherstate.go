package watcherstate

type WatcherState struct {
	Name string
}

var (
	Stopped = WatcherState{"stopped"}

	Starting = WatcherState{"starting"}

	Started = WatcherState{"started"}

	Stopping = WatcherState{"stopping"}
)

func (w WatcherState) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *WatcherState) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (w WatcherState) String() string { _ = "STUB: not implemented"; return "" }
