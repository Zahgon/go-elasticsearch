package types

type UsagePhase struct {
	Actions []string `json:"actions"`
	MinAge  int64    `json:"min_age"`
}

func (s *UsagePhase) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewUsagePhase() *UsagePhase { _ = "STUB: not implemented"; return nil }
