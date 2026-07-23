package types

type AliasDefinition struct {
	Filter *Query `json:"filter,omitempty"`

	IndexRouting *string `json:"index_routing,omitempty"`

	IsHidden *bool `json:"is_hidden,omitempty"`

	IsWriteIndex *bool `json:"is_write_index,omitempty"`

	Routing *string `json:"routing,omitempty"`

	SearchRouting *string `json:"search_routing,omitempty"`
}

func (s *AliasDefinition) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewAliasDefinition() *AliasDefinition { _ = "STUB: not implemented"; return nil }

type AliasDefinitionVariant interface {
	AliasDefinitionCaster() *AliasDefinition
}

func (s *AliasDefinition) AliasDefinitionCaster() *AliasDefinition {
	_ = "STUB: not implemented"
	return nil
}
