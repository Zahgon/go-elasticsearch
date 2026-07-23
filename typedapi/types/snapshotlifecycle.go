package types

type SnapshotLifecycle struct {
	InProgress  *InProgress `json:"in_progress,omitempty"`
	LastFailure *Invocation `json:"last_failure,omitempty"`
	LastSuccess *Invocation `json:"last_success,omitempty"`

	ModifiedDate       DateTime `json:"modified_date,omitempty"`
	ModifiedDateMillis int64    `json:"modified_date_millis"`

	NextExecution       DateTime   `json:"next_execution,omitempty"`
	NextExecutionMillis int64      `json:"next_execution_millis"`
	Policy              SLMPolicy  `json:"policy"`
	Stats               Statistics `json:"stats"`

	Version int64 `json:"version"`
}

func (s *SnapshotLifecycle) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSnapshotLifecycle() *SnapshotLifecycle { _ = "STUB: not implemented"; return nil }
