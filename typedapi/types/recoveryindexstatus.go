package types

type RecoveryIndexStatus struct {
	Bytes                      *RecoveryBytes `json:"bytes,omitempty"`
	Files                      RecoveryFiles  `json:"files"`
	Size                       RecoveryBytes  `json:"size"`
	SourceThrottleTime         Duration       `json:"source_throttle_time,omitempty"`
	SourceThrottleTimeInMillis int64          `json:"source_throttle_time_in_millis"`
	TargetThrottleTime         Duration       `json:"target_throttle_time,omitempty"`
	TargetThrottleTimeInMillis int64          `json:"target_throttle_time_in_millis"`
	TotalTime                  Duration       `json:"total_time,omitempty"`
	TotalTimeInMillis          int64          `json:"total_time_in_millis"`
}

func (s *RecoveryIndexStatus) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewRecoveryIndexStatus() *RecoveryIndexStatus { _ = "STUB: not implemented"; return nil }
