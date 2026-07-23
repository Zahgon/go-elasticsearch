package types

type ModelSnapshot struct {
	Description *string `json:"description,omitempty"`

	JobId string `json:"job_id"`

	LatestRecordTimeStamp *int `json:"latest_record_time_stamp,omitempty"`

	LatestResultTimeStamp *int `json:"latest_result_time_stamp,omitempty"`

	MinVersion string `json:"min_version"`

	ModelSizeStats *ModelSizeStats `json:"model_size_stats,omitempty"`

	Retain bool `json:"retain"`

	SnapshotDocCount int64 `json:"snapshot_doc_count"`

	SnapshotId string `json:"snapshot_id"`

	Timestamp int64 `json:"timestamp"`
}

func (s *ModelSnapshot) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewModelSnapshot() *ModelSnapshot { _ = "STUB: not implemented"; return nil }
