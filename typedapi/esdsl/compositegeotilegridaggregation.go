package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/missingorder"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sortorder"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/valuetype"
)

type _compositeGeoTileGridAggregation struct {
	v *types.CompositeGeoTileGridAggregation
}

func NewCompositeGeoTileGridAggregation() *_compositeGeoTileGridAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_compositeGeoTileGridAggregation) Bounds(geobounds types.GeoBoundsVariant) *_compositeGeoTileGridAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_compositeGeoTileGridAggregation) Precision(precision int) *_compositeGeoTileGridAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_compositeGeoTileGridAggregation) Field(field string) *_compositeGeoTileGridAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_compositeGeoTileGridAggregation) MissingBucket(missingbucket bool) *_compositeGeoTileGridAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_compositeGeoTileGridAggregation) MissingOrder(missingorder missingorder.MissingOrder) *_compositeGeoTileGridAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_compositeGeoTileGridAggregation) Order(order sortorder.SortOrder) *_compositeGeoTileGridAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_compositeGeoTileGridAggregation) Script(script types.ScriptVariant) *_compositeGeoTileGridAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_compositeGeoTileGridAggregation) ValueType(valuetype valuetype.ValueType) *_compositeGeoTileGridAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_compositeGeoTileGridAggregation) CompositeAggregationSourceCaster() *types.CompositeAggregationSource {
	_ = "STUB: not implemented"
	return nil
}

func (s *_compositeGeoTileGridAggregation) CompositeGeoTileGridAggregationCaster() *types.CompositeGeoTileGridAggregation {
	_ = "STUB: not implemented"
	return nil
}
