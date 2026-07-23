package types

type RecoveryRecord struct {
	Bytes *string `json:"bytes,omitempty"`

	BytesPercent Percentage `json:"bytes_percent,omitempty"`

	BytesRecovered *string `json:"bytes_recovered,omitempty"`

	BytesTotal *string `json:"bytes_total,omitempty"`

	Files *string `json:"files,omitempty"`

	FilesPercent Percentage `json:"files_percent,omitempty"`

	FilesRecovered *string `json:"files_recovered,omitempty"`

	FilesTotal *string `json:"files_total,omitempty"`

	Index *string `json:"index,omitempty"`

	Repository *string `json:"repository,omitempty"`

	Shard *string `json:"shard,omitempty"`

	Snapshot *string `json:"snapshot,omitempty"`

	SourceHost *string `json:"source_host,omitempty"`

	SourceNode *string `json:"source_node,omitempty"`

	Stage *string `json:"stage,omitempty"`

	StartTime DateTime `json:"start_time,omitempty"`

	StartTimeMillis *int64 `json:"start_time_millis,omitempty"`

	StopTime DateTime `json:"stop_time,omitempty"`

	StopTimeMillis *int64 `json:"stop_time_millis,omitempty"`

	TargetHost *string `json:"target_host,omitempty"`

	TargetNode *string `json:"target_node,omitempty"`

	Time Duration `json:"time,omitempty"`

	TranslogOps *string `json:"translog_ops,omitempty"`

	TranslogOpsPercent Percentage `json:"translog_ops_percent,omitempty"`

	TranslogOpsRecovered *string `json:"translog_ops_recovered,omitempty"`

	Type *string `json:"type,omitempty"`
}

func (s *RecoveryRecord) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRecoveryRecord() *RecoveryRecord { _ = "STUB: not implemented"; return nil }
