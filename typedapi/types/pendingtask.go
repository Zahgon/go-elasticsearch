package types

type PendingTask struct {
	Executing bool `json:"executing"`

	InsertOrder int `json:"insert_order"`

	Priority string `json:"priority"`

	Source string `json:"source"`

	TimeInQueue Duration `json:"time_in_queue,omitempty"`

	TimeInQueueMillis int64 `json:"time_in_queue_millis"`
}

func (s *PendingTask) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewPendingTask() *PendingTask { _ = "STUB: not implemented"; return nil }
