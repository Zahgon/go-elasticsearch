package types

type GeoTileGridAggregation struct {
	Bounds GeoBounds `json:"bounds,omitempty"`

	Field *string `json:"field,omitempty"`

	Precision *int `json:"precision,omitempty"`

	ShardSize *int `json:"shard_size,omitempty"`

	Size *int `json:"size,omitempty"`
}

func (s *GeoTileGridAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewGeoTileGridAggregation() *GeoTileGridAggregation { _ = "STUB: not implemented"; return nil }

type GeoTileGridAggregationVariant interface {
	GeoTileGridAggregationCaster() *GeoTileGridAggregation
}

func (s *GeoTileGridAggregation) GeoTileGridAggregationCaster() *GeoTileGridAggregation {
	_ = "STUB: not implemented"
	return nil
}
