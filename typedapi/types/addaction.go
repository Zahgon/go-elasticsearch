package types

type AddAction struct {
	Alias *string `json:"alias,omitempty"`

	Aliases []string `json:"aliases,omitempty"`

	Filter *Query `json:"filter,omitempty"`

	Index *string `json:"index,omitempty"`

	IndexRouting *string `json:"index_routing,omitempty"`

	Indices []string `json:"indices,omitempty"`

	IsHidden *bool `json:"is_hidden,omitempty"`

	IsWriteIndex *bool `json:"is_write_index,omitempty"`

	MustExist *bool `json:"must_exist,omitempty"`

	Routing *string `json:"routing,omitempty"`

	SearchRouting *string `json:"search_routing,omitempty"`
}

func (s *AddAction) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewAddAction() *AddAction { _ = "STUB: not implemented"; return nil }

type AddActionVariant interface {
	AddActionCaster() *AddAction
}

func (s *AddAction) AddActionCaster() *AddAction { _ = "STUB: not implemented"; return nil }
