package types

type IndexSettingBlocks struct {
	Metadata            Stringifiedboolean `json:"metadata,omitempty"`
	Read                Stringifiedboolean `json:"read,omitempty"`
	ReadOnly            Stringifiedboolean `json:"read_only,omitempty"`
	ReadOnlyAllowDelete Stringifiedboolean `json:"read_only_allow_delete,omitempty"`
	Write               Stringifiedboolean `json:"write,omitempty"`
}

func (s *IndexSettingBlocks) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewIndexSettingBlocks() *IndexSettingBlocks { _ = "STUB: not implemented"; return nil }

type IndexSettingBlocksVariant interface {
	IndexSettingBlocksCaster() *IndexSettingBlocks
}

func (s *IndexSettingBlocks) IndexSettingBlocksCaster() *IndexSettingBlocks {
	_ = "STUB: not implemented"
	return nil
}
