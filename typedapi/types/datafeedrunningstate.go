package types

type DatafeedRunningState struct {
	RealTimeConfigured bool `json:"real_time_configured"`

	RealTimeRunning bool `json:"real_time_running"`

	SearchInterval *RunningStateSearchInterval `json:"search_interval,omitempty"`
}

func (s *DatafeedRunningState) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDatafeedRunningState() *DatafeedRunningState { _ = "STUB: not implemented"; return nil }
