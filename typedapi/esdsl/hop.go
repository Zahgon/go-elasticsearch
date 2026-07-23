package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _hop struct {
	v *types.Hop
}

func NewHop() *_hop { _ = "STUB: not implemented"; return nil }

func (s *_hop) Connections(connections types.HopVariant) *_hop {
	_ = "STUB: not implemented"
	return nil
}

func (s *_hop) Query(query types.QueryVariant) *_hop { _ = "STUB: not implemented"; return nil }

func (s *_hop) Vertices(vertices ...types.VertexDefinitionVariant) *_hop {
	_ = "STUB: not implemented"
	return nil
}

func (s *_hop) VerticesValues(verticesvalues []types.VertexDefinition) *_hop {
	_ = "STUB: not implemented"
	return nil
}

func (s *_hop) HopCaster() *types.Hop { _ = "STUB: not implemented"; return nil }
