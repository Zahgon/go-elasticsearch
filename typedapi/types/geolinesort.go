package types

type GeoLineSort struct {
	Field string `json:"field"`
}

func (s *GeoLineSort) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewGeoLineSort() *GeoLineSort { _ = "STUB: not implemented"; return nil }

type GeoLineSortVariant interface {
	GeoLineSortCaster() *GeoLineSort
}

func (s *GeoLineSort) GeoLineSortCaster() *GeoLineSort { _ = "STUB: not implemented"; return nil }
