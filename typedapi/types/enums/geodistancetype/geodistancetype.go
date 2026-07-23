package geodistancetype

type GeoDistanceType struct {
	Name string
}

var (
	Arc = GeoDistanceType{"arc"}

	Plane = GeoDistanceType{"plane"}
)

func (g GeoDistanceType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *GeoDistanceType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (g GeoDistanceType) String() string { _ = "STUB: not implemented"; return "" }
