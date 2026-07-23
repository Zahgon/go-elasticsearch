package types

type PerRepositoryStats struct {
	CurrentCounts         RepositoryStatsCurrentCounts `json:"current_counts"`
	OldestStartTime       *string                      `json:"oldest_start_time,omitempty"`
	OldestStartTimeMillis int64                        `json:"oldest_start_time_millis"`
	Type                  string                       `json:"type"`
}

func (s *PerRepositoryStats) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewPerRepositoryStats() *PerRepositoryStats { _ = "STUB: not implemented"; return nil }
