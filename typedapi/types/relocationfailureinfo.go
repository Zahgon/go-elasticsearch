package types

type RelocationFailureInfo struct {
	FailedAttempts int `json:"failed_attempts"`
}

func (s *RelocationFailureInfo) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewRelocationFailureInfo() *RelocationFailureInfo { _ = "STUB: not implemented"; return nil }
