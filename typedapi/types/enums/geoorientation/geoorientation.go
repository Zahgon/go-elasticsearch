package geoorientation

type GeoOrientation struct {
	Name string
}

var (
	Right = GeoOrientation{"right"}

	Left = GeoOrientation{"left"}
)

func (g GeoOrientation) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *GeoOrientation) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (g GeoOrientation) String() string { _ = "STUB: not implemented"; return "" }
