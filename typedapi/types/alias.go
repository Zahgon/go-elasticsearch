package types

type Alias struct {
	Filter *Query `json:"filter,omitempty"`

	IndexRouting *string `json:"index_routing,omitempty"`

	IsHidden *bool `json:"is_hidden,omitempty"`

	IsWriteIndex *bool `json:"is_write_index,omitempty"`

	Routing *string `json:"routing,omitempty"`

	SearchRouting *string `json:"search_routing,omitempty"`
}

func (s *Alias) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewAlias() *Alias { _ = "STUB: not implemented"; return nil }

type AliasVariant interface {
	AliasCaster() *Alias
}

func (s *Alias) AliasCaster() *Alias { _ = "STUB: not implemented"; return nil }
