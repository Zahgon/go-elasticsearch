package types

type CompositeAggregationSource struct {
	DateHistogram *CompositeDateHistogramAggregation `json:"date_histogram,omitempty"`

	GeotileGrid *CompositeGeoTileGridAggregation `json:"geotile_grid,omitempty"`

	Histogram *CompositeHistogramAggregation `json:"histogram,omitempty"`

	Terms *CompositeTermsAggregation `json:"terms,omitempty"`
}

func NewCompositeAggregationSource() *CompositeAggregationSource {
	_ = "STUB: not implemented"
	return nil
}

type CompositeAggregationSourceVariant interface {
	CompositeAggregationSourceCaster() *CompositeAggregationSource
}

func (s *CompositeAggregationSource) CompositeAggregationSourceCaster() *CompositeAggregationSource {
	_ = "STUB: not implemented"
	return nil
}
