package types

type Watcher struct {
	Available bool           `json:"available"`
	Count     Counter        `json:"count"`
	Enabled   bool           `json:"enabled"`
	Execution WatcherActions `json:"execution"`
	Watch     WatcherWatch   `json:"watch"`
}

func (s *Watcher) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewWatcher() *Watcher { _ = "STUB: not implemented"; return nil }
