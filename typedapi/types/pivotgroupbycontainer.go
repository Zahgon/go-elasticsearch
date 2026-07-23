package types

type PivotGroupByContainer struct {
	DateHistogram *DateHistogramAggregation `json:"date_histogram,omitempty"`
	GeotileGrid   *GeoTileGridAggregation   `json:"geotile_grid,omitempty"`
	Histogram     *HistogramAggregation     `json:"histogram,omitempty"`
	Terms         *TermsAggregation         `json:"terms,omitempty"`
}

func NewPivotGroupByContainer() *PivotGroupByContainer { _ = "STUB: not implemented"; return nil }

type PivotGroupByContainerVariant interface {
	PivotGroupByContainerCaster() *PivotGroupByContainer
}

func (s *PivotGroupByContainer) PivotGroupByContainerCaster() *PivotGroupByContainer {
	_ = "STUB: not implemented"
	return nil
}
