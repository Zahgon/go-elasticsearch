package types

type GetStats struct {
	Current             int64    `json:"current"`
	ExistsTime          Duration `json:"exists_time,omitempty"`
	ExistsTimeInMillis  int64    `json:"exists_time_in_millis"`
	ExistsTotal         int64    `json:"exists_total"`
	MissingTime         Duration `json:"missing_time,omitempty"`
	MissingTimeInMillis int64    `json:"missing_time_in_millis"`
	MissingTotal        int64    `json:"missing_total"`
	Time                Duration `json:"time,omitempty"`
	TimeInMillis        int64    `json:"time_in_millis"`
	Total               int64    `json:"total"`
}

func (s *GetStats) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewGetStats() *GetStats { _ = "STUB: not implemented"; return nil }
