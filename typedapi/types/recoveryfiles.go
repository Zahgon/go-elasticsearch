package types

type RecoveryFiles struct {
	Details   []FileDetails `json:"details,omitempty"`
	Percent   Percentage    `json:"percent"`
	Recovered int64         `json:"recovered"`
	Reused    int64         `json:"reused"`
	Total     int64         `json:"total"`
}

func (s *RecoveryFiles) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRecoveryFiles() *RecoveryFiles { _ = "STUB: not implemented"; return nil }
