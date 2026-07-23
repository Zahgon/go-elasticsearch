package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/geodistancetype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/geovalidationmethod"
)

type GeoDistanceQuery struct {
	Boost *float32 `json:"boost,omitempty"`

	Distance string `json:"distance"`

	DistanceType     *geodistancetype.GeoDistanceType `json:"distance_type,omitempty"`
	GeoDistanceQuery map[string]GeoLocation           `json:"-"`

	IgnoreUnmapped *bool   `json:"ignore_unmapped,omitempty"`
	QueryName_     *string `json:"_name,omitempty"`

	ValidationMethod *geovalidationmethod.GeoValidationMethod `json:"validation_method,omitempty"`
}

func (s *GeoDistanceQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s GeoDistanceQuery) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewGeoDistanceQuery() *GeoDistanceQuery { _ = "STUB: not implemented"; return nil }

type GeoDistanceQueryVariant interface {
	GeoDistanceQueryCaster() *GeoDistanceQuery
}

func (s *GeoDistanceQuery) GeoDistanceQueryCaster() *GeoDistanceQuery {
	_ = "STUB: not implemented"
	return nil
}
