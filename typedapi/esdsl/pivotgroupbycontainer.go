package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _pivotGroupByContainer struct {
	v *types.PivotGroupByContainer
}

func NewPivotGroupByContainer() *_pivotGroupByContainer { _ = "STUB: not implemented"; return nil }

func (s *_pivotGroupByContainer) DateHistogram(datehistogram types.DateHistogramAggregationVariant) *_pivotGroupByContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pivotGroupByContainer) GeotileGrid(geotilegrid types.GeoTileGridAggregationVariant) *_pivotGroupByContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pivotGroupByContainer) Histogram(histogram types.HistogramAggregationVariant) *_pivotGroupByContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pivotGroupByContainer) Terms(terms types.TermsAggregationVariant) *_pivotGroupByContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pivotGroupByContainer) PivotGroupByContainerCaster() *types.PivotGroupByContainer {
	_ = "STUB: not implemented"
	return nil
}
