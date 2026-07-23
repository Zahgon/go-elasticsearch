package types

type VerifyIndex struct {
	CheckIndexTime         Duration `json:"check_index_time,omitempty"`
	CheckIndexTimeInMillis int64    `json:"check_index_time_in_millis"`
	TotalTime              Duration `json:"total_time,omitempty"`
	TotalTimeInMillis      int64    `json:"total_time_in_millis"`
}

func (s *VerifyIndex) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewVerifyIndex() *VerifyIndex { _ = "STUB: not implemented"; return nil }
