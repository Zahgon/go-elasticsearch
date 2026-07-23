package types

type RolloverConditions struct {
	MaxAge                   Duration `json:"max_age,omitempty"`
	MaxAgeMillis             *int64   `json:"max_age_millis,omitempty"`
	MaxDocs                  *int64   `json:"max_docs,omitempty"`
	MaxPrimaryShardDocs      *int64   `json:"max_primary_shard_docs,omitempty"`
	MaxPrimaryShardSize      ByteSize `json:"max_primary_shard_size,omitempty"`
	MaxPrimaryShardSizeBytes *int64   `json:"max_primary_shard_size_bytes,omitempty"`

	MaxSize                  ByteSize `json:"max_size,omitempty"`
	MaxSizeBytes             *int64   `json:"max_size_bytes,omitempty"`
	MinAge                   Duration `json:"min_age,omitempty"`
	MinDocs                  *int64   `json:"min_docs,omitempty"`
	MinPrimaryShardDocs      *int64   `json:"min_primary_shard_docs,omitempty"`
	MinPrimaryShardSize      ByteSize `json:"min_primary_shard_size,omitempty"`
	MinPrimaryShardSizeBytes *int64   `json:"min_primary_shard_size_bytes,omitempty"`
	MinSize                  ByteSize `json:"min_size,omitempty"`
	MinSizeBytes             *int64   `json:"min_size_bytes,omitempty"`
}

func (s *RolloverConditions) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewRolloverConditions() *RolloverConditions { _ = "STUB: not implemented"; return nil }

type RolloverConditionsVariant interface {
	RolloverConditionsCaster() *RolloverConditions
}

func (s *RolloverConditions) RolloverConditionsCaster() *RolloverConditions {
	_ = "STUB: not implemented"
	return nil
}
