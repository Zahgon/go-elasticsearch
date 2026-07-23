package types

type ShrinkAction struct {
	AllowWriteAfterShrink *bool    `json:"allow_write_after_shrink,omitempty"`
	MaxPrimaryShardSize   ByteSize `json:"max_primary_shard_size,omitempty"`
	NumberOfShards        *int     `json:"number_of_shards,omitempty"`
}

func (s *ShrinkAction) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewShrinkAction() *ShrinkAction { _ = "STUB: not implemented"; return nil }

type ShrinkActionVariant interface {
	ShrinkActionCaster() *ShrinkAction
}

func (s *ShrinkAction) ShrinkActionCaster() *ShrinkAction { _ = "STUB: not implemented"; return nil }
