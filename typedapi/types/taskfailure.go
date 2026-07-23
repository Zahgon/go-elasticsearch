package types

type TaskFailure struct {
	NodeId string     `json:"node_id"`
	Reason ErrorCause `json:"reason"`
	Status string     `json:"status"`
	TaskId int64      `json:"task_id"`
}

func (s *TaskFailure) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewTaskFailure() *TaskFailure { _ = "STUB: not implemented"; return nil }
