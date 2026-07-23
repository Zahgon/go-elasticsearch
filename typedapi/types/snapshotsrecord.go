package types

type SnapshotsRecord struct {
	Duration Duration `json:"duration,omitempty"`

	EndEpoch StringifiedEpochTimeUnitSeconds `json:"end_epoch,omitempty"`

	EndTime *string `json:"end_time,omitempty"`

	FailedShards *string `json:"failed_shards,omitempty"`

	Id *string `json:"id,omitempty"`

	Indices *string `json:"indices,omitempty"`

	Reason *string `json:"reason,omitempty"`

	Repository *string `json:"repository,omitempty"`

	StartEpoch StringifiedEpochTimeUnitSeconds `json:"start_epoch,omitempty"`

	StartTime ScheduleTimeOfDay `json:"start_time,omitempty"`

	Status *string `json:"status,omitempty"`

	SuccessfulShards *string `json:"successful_shards,omitempty"`

	TotalShards *string `json:"total_shards,omitempty"`
}

func (s *SnapshotsRecord) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSnapshotsRecord() *SnapshotsRecord { _ = "STUB: not implemented"; return nil }
