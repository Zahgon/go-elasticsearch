package types

type IndexingStats struct {
	DeleteCurrent        int64                    `json:"delete_current"`
	DeleteTime           Duration                 `json:"delete_time,omitempty"`
	DeleteTimeInMillis   int64                    `json:"delete_time_in_millis"`
	DeleteTotal          int64                    `json:"delete_total"`
	IndexCurrent         int64                    `json:"index_current"`
	IndexFailed          int64                    `json:"index_failed"`
	IndexTime            Duration                 `json:"index_time,omitempty"`
	IndexTimeInMillis    int64                    `json:"index_time_in_millis"`
	IndexTotal           int64                    `json:"index_total"`
	IsThrottled          bool                     `json:"is_throttled"`
	NoopUpdateTotal      int64                    `json:"noop_update_total"`
	PeakWriteLoad        *Float64                 `json:"peak_write_load,omitempty"`
	RecentWriteLoad      *Float64                 `json:"recent_write_load,omitempty"`
	ThrottleTime         Duration                 `json:"throttle_time,omitempty"`
	ThrottleTimeInMillis int64                    `json:"throttle_time_in_millis"`
	Types                map[string]IndexingStats `json:"types,omitempty"`
	WriteLoad            *Float64                 `json:"write_load,omitempty"`
}

func (s *IndexingStats) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewIndexingStats() *IndexingStats { _ = "STUB: not implemented"; return nil }
