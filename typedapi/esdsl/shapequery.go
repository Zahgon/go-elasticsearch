package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _shapeQuery struct {
	v *types.ShapeQuery
}

func NewShapeQuery() *_shapeQuery { _ = "STUB: not implemented"; return nil }

func (s *_shapeQuery) IgnoreUnmapped(ignoreunmapped bool) *_shapeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_shapeQuery) ShapeQuery(shapequery map[string]types.ShapeFieldQuery) *_shapeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_shapeQuery) AddShapeQuery(key string, value types.ShapeFieldQueryVariant) *_shapeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_shapeQuery) Boost(boost float32) *_shapeQuery { _ = "STUB: not implemented"; return nil }

func (s *_shapeQuery) QueryName_(queryname_ string) *_shapeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_shapeQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func (s *_shapeQuery) ShapeQueryCaster() *types.ShapeQuery { _ = "STUB: not implemented"; return nil }
