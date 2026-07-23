package geopointmetrictype

type GeoPointMetricType struct {
	Name string
}

var (
	Gauge = GeoPointMetricType{"gauge"}

	Counter = GeoPointMetricType{"counter"}

	Position = GeoPointMetricType{"position"}
)

func (g GeoPointMetricType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *GeoPointMetricType) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (g GeoPointMetricType) String() string { _ = "STUB: not implemented"; return "" }
