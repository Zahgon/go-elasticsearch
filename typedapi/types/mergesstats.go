package types

type MergesStats struct {
	Current                    int64    `json:"current"`
	CurrentDocs                int64    `json:"current_docs"`
	CurrentSize                *string  `json:"current_size,omitempty"`
	CurrentSizeInBytes         int64    `json:"current_size_in_bytes"`
	Total                      int64    `json:"total"`
	TotalAutoThrottle          *string  `json:"total_auto_throttle,omitempty"`
	TotalAutoThrottleInBytes   int64    `json:"total_auto_throttle_in_bytes"`
	TotalDocs                  int64    `json:"total_docs"`
	TotalSize                  *string  `json:"total_size,omitempty"`
	TotalSizeInBytes           int64    `json:"total_size_in_bytes"`
	TotalStoppedTime           Duration `json:"total_stopped_time,omitempty"`
	TotalStoppedTimeInMillis   int64    `json:"total_stopped_time_in_millis"`
	TotalThrottledTime         Duration `json:"total_throttled_time,omitempty"`
	TotalThrottledTimeInMillis int64    `json:"total_throttled_time_in_millis"`
	TotalTime                  Duration `json:"total_time,omitempty"`
	TotalTimeInMillis          int64    `json:"total_time_in_millis"`
}

func (s *MergesStats) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewMergesStats() *MergesStats { _ = "STUB: not implemented"; return nil }
