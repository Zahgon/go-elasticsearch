package types

type HealthRecord struct {
	ActiveShardsPercent *string `json:"active_shards_percent,omitempty"`

	Cluster *string `json:"cluster,omitempty"`

	Epoch StringifiedEpochTimeUnitSeconds `json:"epoch,omitempty"`

	Init *string `json:"init,omitempty"`

	MaxTaskWaitTime *string `json:"max_task_wait_time,omitempty"`

	NodeData *string `json:"node.data,omitempty"`

	NodeTotal *string `json:"node.total,omitempty"`

	PendingTasks *string `json:"pending_tasks,omitempty"`

	Pri *string `json:"pri,omitempty"`

	Relo *string `json:"relo,omitempty"`

	Shards *string `json:"shards,omitempty"`

	Status *string `json:"status,omitempty"`

	Timestamp *string `json:"timestamp,omitempty"`

	Unassign *string `json:"unassign,omitempty"`

	UnassignPri *string `json:"unassign.pri,omitempty"`
}

func (s *HealthRecord) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewHealthRecord() *HealthRecord { _ = "STUB: not implemented"; return nil }
