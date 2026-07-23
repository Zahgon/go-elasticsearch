package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _vertexDefinition struct {
	v *types.VertexDefinition
}

func NewVertexDefinition() *_vertexDefinition { _ = "STUB: not implemented"; return nil }

func (s *_vertexDefinition) Exclude(excludes ...string) *_vertexDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_vertexDefinition) Field(field string) *_vertexDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_vertexDefinition) Include(includes ...types.VertexIncludeVariant) *_vertexDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_vertexDefinition) IncludeValues(includevalues []types.VertexInclude) *_vertexDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_vertexDefinition) MinDocCount(mindoccount int64) *_vertexDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_vertexDefinition) ShardMinDocCount(shardmindoccount int64) *_vertexDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_vertexDefinition) Size(size int) *_vertexDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_vertexDefinition) VertexDefinitionCaster() *types.VertexDefinition {
	_ = "STUB: not implemented"
	return nil
}
