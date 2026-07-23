package types

type Recording struct {
	CumulativeExecutionCount      *int64   `json:"cumulative_execution_count,omitempty"`
	CumulativeExecutionTime       Duration `json:"cumulative_execution_time,omitempty"`
	CumulativeExecutionTimeMillis *int64   `json:"cumulative_execution_time_millis,omitempty"`
	Name                          *string  `json:"name,omitempty"`
}

func (s *Recording) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRecording() *Recording { _ = "STUB: not implemented"; return nil }
