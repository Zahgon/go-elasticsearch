package types

type CreateFrom struct {
	MappingsOverride *TypeMapping `json:"mappings_override,omitempty"`

	RemoveIndexBlocks *bool `json:"remove_index_blocks,omitempty"`

	SettingsOverride *IndexSettings `json:"settings_override,omitempty"`
}

func (s *CreateFrom) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewCreateFrom() *CreateFrom { _ = "STUB: not implemented"; return nil }

type CreateFromVariant interface {
	CreateFromCaster() *CreateFrom
}

func (s *CreateFrom) CreateFromCaster() *CreateFrom { _ = "STUB: not implemented"; return nil }
