package types

type RecoveryStartStatus struct {
	CheckIndexTime         Duration `json:"check_index_time,omitempty"`
	CheckIndexTimeInMillis int64    `json:"check_index_time_in_millis"`
	TotalTime              Duration `json:"total_time,omitempty"`
	TotalTimeInMillis      int64    `json:"total_time_in_millis"`
}

func (s *RecoveryStartStatus) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewRecoveryStartStatus() *RecoveryStartStatus { _ = "STUB: not implemented"; return nil }
