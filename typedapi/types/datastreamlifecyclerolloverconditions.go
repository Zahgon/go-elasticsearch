package types

type DataStreamLifecycleRolloverConditions struct {
	MaxAge              *string  `json:"max_age,omitempty"`
	MaxDocs             *int64   `json:"max_docs,omitempty"`
	MaxPrimaryShardDocs *int64   `json:"max_primary_shard_docs,omitempty"`
	MaxPrimaryShardSize ByteSize `json:"max_primary_shard_size,omitempty"`
	MaxSize             ByteSize `json:"max_size,omitempty"`
	MinAge              Duration `json:"min_age,omitempty"`
	MinDocs             *int64   `json:"min_docs,omitempty"`
	MinPrimaryShardDocs *int64   `json:"min_primary_shard_docs,omitempty"`
	MinPrimaryShardSize ByteSize `json:"min_primary_shard_size,omitempty"`
	MinSize             ByteSize `json:"min_size,omitempty"`
}

func (s *DataStreamLifecycleRolloverConditions) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDataStreamLifecycleRolloverConditions() *DataStreamLifecycleRolloverConditions {
	_ = "STUB: not implemented"
	return nil
}
