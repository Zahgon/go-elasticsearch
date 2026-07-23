package types

type TasksRecord struct {
	Action *string `json:"action,omitempty"`

	Description *string `json:"description,omitempty"`

	Id *string `json:"id,omitempty"`

	Ip *string `json:"ip,omitempty"`

	Node *string `json:"node,omitempty"`

	NodeId *string `json:"node_id,omitempty"`

	ParentTaskId *string `json:"parent_task_id,omitempty"`

	Port *string `json:"port,omitempty"`

	RunningTime *string `json:"running_time,omitempty"`

	RunningTimeNs *string `json:"running_time_ns,omitempty"`

	StartTime *string `json:"start_time,omitempty"`

	TaskId *string `json:"task_id,omitempty"`

	Timestamp *string `json:"timestamp,omitempty"`

	Type *string `json:"type,omitempty"`

	Version *string `json:"version,omitempty"`

	XOpaqueId *string `json:"x_opaque_id,omitempty"`
}

func (s *TasksRecord) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewTasksRecord() *TasksRecord { _ = "STUB: not implemented"; return nil }
