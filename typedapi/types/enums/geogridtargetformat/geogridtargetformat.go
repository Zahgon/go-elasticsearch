package geogridtargetformat

type GeoGridTargetFormat struct {
	Name string
}

var (
	Geojson = GeoGridTargetFormat{"geojson"}

	Wkt = GeoGridTargetFormat{"wkt"}
)

func (g GeoGridTargetFormat) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *GeoGridTargetFormat) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (g GeoGridTargetFormat) String() string { _ = "STUB: not implemented"; return "" }
