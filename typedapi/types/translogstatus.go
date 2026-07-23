package types

type TranslogStatus struct {
	Percent           Percentage `json:"percent"`
	Recovered         int64      `json:"recovered"`
	Total             int64      `json:"total"`
	TotalOnStart      int64      `json:"total_on_start"`
	TotalTime         Duration   `json:"total_time,omitempty"`
	TotalTimeInMillis int64      `json:"total_time_in_millis"`
}

func (s *TranslogStatus) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewTranslogStatus() *TranslogStatus { _ = "STUB: not implemented"; return nil }
