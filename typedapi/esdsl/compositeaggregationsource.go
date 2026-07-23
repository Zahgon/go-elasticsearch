package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _compositeAggregationSource struct {
	v *types.CompositeAggregationSource
}

func NewCompositeAggregationSource() *_compositeAggregationSource {
	_ = "STUB: not implemented"
	return nil
}

func (s *_compositeAggregationSource) DateHistogram(datehistogram types.CompositeDateHistogramAggregationVariant) *_compositeAggregationSource {
	_ = "STUB: not implemented"
	return nil
}

func (s *_compositeAggregationSource) GeotileGrid(geotilegrid types.CompositeGeoTileGridAggregationVariant) *_compositeAggregationSource {
	_ = "STUB: not implemented"
	return nil
}

func (s *_compositeAggregationSource) Histogram(histogram types.CompositeHistogramAggregationVariant) *_compositeAggregationSource {
	_ = "STUB: not implemented"
	return nil
}

func (s *_compositeAggregationSource) Terms(terms types.CompositeTermsAggregationVariant) *_compositeAggregationSource {
	_ = "STUB: not implemented"
	return nil
}

func (s *_compositeAggregationSource) CompositeAggregationSourceCaster() *types.CompositeAggregationSource {
	_ = "STUB: not implemented"
	return nil
}
