package types

type PendingTasksRecord struct {
	InsertOrder *string `json:"insertOrder,omitempty"`

	Priority *string `json:"priority,omitempty"`

	Source *string `json:"source,omitempty"`

	TimeInQueue *string `json:"timeInQueue,omitempty"`
}

func (s *PendingTasksRecord) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewPendingTasksRecord() *PendingTasksRecord { _ = "STUB: not implemented"; return nil }
