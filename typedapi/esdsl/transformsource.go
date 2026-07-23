package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _transformSource struct {
	v *types.TransformSource
}

func NewTransformSource() *_transformSource { _ = "STUB: not implemented"; return nil }

func (s *_transformSource) Index(indices ...string) *_transformSource {
	_ = "STUB: not implemented"
	return nil
}

func (s *_transformSource) ProjectRouting(projectrouting string) *_transformSource {
	_ = "STUB: not implemented"
	return nil
}

func (s *_transformSource) Query(query types.QueryVariant) *_transformSource {
	_ = "STUB: not implemented"
	return nil
}

func (s *_transformSource) RuntimeMappings(runtimefields types.RuntimeFieldsVariant) *_transformSource {
	_ = "STUB: not implemented"
	return nil
}

func (s *_transformSource) TransformSourceCaster() *types.TransformSource {
	_ = "STUB: not implemented"
	return nil
}
