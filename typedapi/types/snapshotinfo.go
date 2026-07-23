package types

type SnapshotInfo struct {
	DataStreams        []string                `json:"data_streams"`
	Duration           Duration                `json:"duration,omitempty"`
	DurationInMillis   *int64                  `json:"duration_in_millis,omitempty"`
	EndTime            DateTime                `json:"end_time,omitempty"`
	EndTimeInMillis    *int64                  `json:"end_time_in_millis,omitempty"`
	Failures           []SnapshotShardFailure  `json:"failures,omitempty"`
	FeatureStates      []InfoFeatureState      `json:"feature_states,omitempty"`
	IncludeGlobalState *bool                   `json:"include_global_state,omitempty"`
	IndexDetails       map[string]IndexDetails `json:"index_details,omitempty"`
	Indices            []string                `json:"indices,omitempty"`
	Metadata           Metadata                `json:"metadata,omitempty"`
	Reason             *string                 `json:"reason,omitempty"`
	Repository         *string                 `json:"repository,omitempty"`
	Shards             *ShardStatistics        `json:"shards,omitempty"`
	Snapshot           string                  `json:"snapshot"`
	StartTime          DateTime                `json:"start_time,omitempty"`
	StartTimeInMillis  *int64                  `json:"start_time_in_millis,omitempty"`
	State              *string                 `json:"state,omitempty"`
	Uuid               string                  `json:"uuid"`
	Version            *string                 `json:"version,omitempty"`
	VersionId          *int64                  `json:"version_id,omitempty"`
}

func (s *SnapshotInfo) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSnapshotInfo() *SnapshotInfo { _ = "STUB: not implemented"; return nil }
