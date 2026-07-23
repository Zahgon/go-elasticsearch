package types

type VertexDefinition struct {
	Exclude []string `json:"exclude,omitempty"`

	Field string `json:"field"`

	Include []VertexInclude `json:"include,omitempty"`

	MinDocCount *int64 `json:"min_doc_count,omitempty"`

	ShardMinDocCount *int64 `json:"shard_min_doc_count,omitempty"`

	Size *int `json:"size,omitempty"`
}

func (s *VertexDefinition) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewVertexDefinition() *VertexDefinition { _ = "STUB: not implemented"; return nil }

type VertexDefinitionVariant interface {
	VertexDefinitionCaster() *VertexDefinition
}

func (s *VertexDefinition) VertexDefinitionCaster() *VertexDefinition {
	_ = "STUB: not implemented"
	return nil
}
