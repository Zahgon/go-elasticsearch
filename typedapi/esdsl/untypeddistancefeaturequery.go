package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _untypedDistanceFeatureQuery struct {
	v *types.UntypedDistanceFeatureQuery
}

func NewUntypedDistanceFeatureQuery(origin json.RawMessage, pivot json.RawMessage) *_untypedDistanceFeatureQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_untypedDistanceFeatureQuery) Boost(boost float32) *_untypedDistanceFeatureQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_untypedDistanceFeatureQuery) Field(field string) *_untypedDistanceFeatureQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_untypedDistanceFeatureQuery) Origin(origin json.RawMessage) *_untypedDistanceFeatureQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_untypedDistanceFeatureQuery) Pivot(pivot json.RawMessage) *_untypedDistanceFeatureQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_untypedDistanceFeatureQuery) QueryName_(queryname_ string) *_untypedDistanceFeatureQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_untypedDistanceFeatureQuery) QueryCaster() *types.Query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_untypedDistanceFeatureQuery) UntypedDistanceFeatureQueryCaster() *types.UntypedDistanceFeatureQuery {
	_ = "STUB: not implemented"
	return nil
}
