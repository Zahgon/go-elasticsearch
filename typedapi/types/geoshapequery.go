package types

type GeoShapeQuery struct {
	Boost         *float32                      `json:"boost,omitempty"`
	GeoShapeQuery map[string]GeoShapeFieldQuery `json:"-"`

	IgnoreUnmapped *bool   `json:"ignore_unmapped,omitempty"`
	QueryName_     *string `json:"_name,omitempty"`
}

func (s *GeoShapeQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s GeoShapeQuery) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewGeoShapeQuery() *GeoShapeQuery { _ = "STUB: not implemented"; return nil }

type GeoShapeQueryVariant interface {
	GeoShapeQueryCaster() *GeoShapeQuery
}

func (s *GeoShapeQuery) GeoShapeQueryCaster() *GeoShapeQuery { _ = "STUB: not implemented"; return nil }
