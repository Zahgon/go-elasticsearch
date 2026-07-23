package types

import (
	"encoding/json"
)

type GeoLineAggregate struct {
	Geometry   GeoLine         `json:"geometry"`
	Meta       Metadata        `json:"meta,omitempty"`
	Properties json.RawMessage `json:"properties,omitempty"`
	Type       string          `json:"type"`
}

func (s *GeoLineAggregate) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewGeoLineAggregate() *GeoLineAggregate { _ = "STUB: not implemented"; return nil }
