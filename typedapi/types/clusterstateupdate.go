package types

type ClusterStateUpdate struct {
	CommitTime Duration `json:"commit_time,omitempty"`

	CommitTimeMillis *int64 `json:"commit_time_millis,omitempty"`

	CompletionTime Duration `json:"completion_time,omitempty"`

	CompletionTimeMillis *int64 `json:"completion_time_millis,omitempty"`

	ComputationTime Duration `json:"computation_time,omitempty"`

	ComputationTimeMillis *int64 `json:"computation_time_millis,omitempty"`

	ContextConstructionTime Duration `json:"context_construction_time,omitempty"`

	ContextConstructionTimeMillis *int64 `json:"context_construction_time_millis,omitempty"`

	Count int64 `json:"count"`

	MasterApplyTime Duration `json:"master_apply_time,omitempty"`

	MasterApplyTimeMillis *int64 `json:"master_apply_time_millis,omitempty"`

	NotificationTime Duration `json:"notification_time,omitempty"`

	NotificationTimeMillis *int64 `json:"notification_time_millis,omitempty"`

	PublicationTime Duration `json:"publication_time,omitempty"`

	PublicationTimeMillis *int64 `json:"publication_time_millis,omitempty"`
}

func (s *ClusterStateUpdate) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewClusterStateUpdate() *ClusterStateUpdate { _ = "STUB: not implemented"; return nil }
