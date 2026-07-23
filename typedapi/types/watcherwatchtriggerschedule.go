package types

type WatcherWatchTriggerSchedule struct {
	Active int64   `json:"active"`
	All_   Counter `json:"_all"`
	Cron   Counter `json:"cron"`
	Total  int64   `json:"total"`
}

func (s *WatcherWatchTriggerSchedule) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewWatcherWatchTriggerSchedule() *WatcherWatchTriggerSchedule {
	_ = "STUB: not implemented"
	return nil
}
