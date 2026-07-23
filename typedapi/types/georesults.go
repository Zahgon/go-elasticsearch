package types

type GeoResults struct {
	ActualPoint *string `json:"actual_point,omitempty"`

	TypicalPoint *string `json:"typical_point,omitempty"`
}

func (s *GeoResults) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewGeoResults() *GeoResults { _ = "STUB: not implemented"; return nil }
