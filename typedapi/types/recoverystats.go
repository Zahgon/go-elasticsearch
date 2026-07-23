package types

type RecoveryStats struct {
	CurrentAsSource      int64    `json:"current_as_source"`
	CurrentAsTarget      int64    `json:"current_as_target"`
	ThrottleTime         Duration `json:"throttle_time,omitempty"`
	ThrottleTimeInMillis int64    `json:"throttle_time_in_millis"`
}

func (s *RecoveryStats) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRecoveryStats() *RecoveryStats { _ = "STUB: not implemented"; return nil }
