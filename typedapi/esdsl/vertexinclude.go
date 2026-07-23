package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _vertexInclude struct {
	v *types.VertexInclude
}

func NewVertexInclude(term string) *_vertexInclude { _ = "STUB: not implemented"; return nil }

func (s *_vertexInclude) Boost(boost types.Float64) *_vertexInclude {
	_ = "STUB: not implemented"
	return nil
}

func (s *_vertexInclude) Term(term string) *_vertexInclude { _ = "STUB: not implemented"; return nil }

func (s *_vertexInclude) VertexIncludeCaster() *types.VertexInclude {
	_ = "STUB: not implemented"
	return nil
}
