package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/missingorder"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sortorder"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/valuetype"
)

type CompositeGeoTileGridAggregation struct {
	Bounds GeoBounds `json:"bounds,omitempty"`

	Field         *string                    `json:"field,omitempty"`
	MissingBucket *bool                      `json:"missing_bucket,omitempty"`
	MissingOrder  *missingorder.MissingOrder `json:"missing_order,omitempty"`
	Order         *sortorder.SortOrder       `json:"order,omitempty"`
	Precision     *int                       `json:"precision,omitempty"`

	Script    *Script              `json:"script,omitempty"`
	ValueType *valuetype.ValueType `json:"value_type,omitempty"`
}

func (s *CompositeGeoTileGridAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewCompositeGeoTileGridAggregation() *CompositeGeoTileGridAggregation {
	_ = "STUB: not implemented"
	return nil
}

type CompositeGeoTileGridAggregationVariant interface {
	CompositeGeoTileGridAggregationCaster() *CompositeGeoTileGridAggregation
}

func (s *CompositeGeoTileGridAggregation) CompositeGeoTileGridAggregationCaster() *CompositeGeoTileGridAggregation {
	_ = "STUB: not implemented"
	return nil
}
