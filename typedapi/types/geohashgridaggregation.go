package types

type GeoHashGridAggregation struct {
	Bounds GeoBounds `json:"bounds,omitempty"`

	Field *string `json:"field,omitempty"`

	Precision GeoHashPrecision `json:"precision,omitempty"`

	ShardSize *int `json:"shard_size,omitempty"`

	Size *int `json:"size,omitempty"`
}

func (s *GeoHashGridAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewGeoHashGridAggregation() *GeoHashGridAggregation { _ = "STUB: not implemented"; return nil }

type GeoHashGridAggregationVariant interface {
	GeoHashGridAggregationCaster() *GeoHashGridAggregation
}

func (s *GeoHashGridAggregation) GeoHashGridAggregationCaster() *GeoHashGridAggregation {
	_ = "STUB: not implemented"
	return nil
}
