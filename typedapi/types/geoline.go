package types

type GeoLine struct {
	Coordinates [][]Float64 `json:"coordinates"`

	Type string `json:"type"`
}

func (s *GeoLine) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewGeoLine() *GeoLine { _ = "STUB: not implemented"; return nil }
