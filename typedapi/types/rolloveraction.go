package types

type RolloverAction struct {
	MaxAge              Duration `json:"max_age,omitempty"`
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

func (s *RolloverAction) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRolloverAction() *RolloverAction { _ = "STUB: not implemented"; return nil }

type RolloverActionVariant interface {
	RolloverActionCaster() *RolloverAction
}

func (s *RolloverAction) RolloverActionCaster() *RolloverAction {
	_ = "STUB: not implemented"
	return nil
}
