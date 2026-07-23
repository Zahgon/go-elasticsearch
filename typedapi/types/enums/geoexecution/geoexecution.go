package geoexecution

type GeoExecution struct {
	Name string
}

var (
	Memory = GeoExecution{"memory"}

	Indexed = GeoExecution{"indexed"}
)

func (g GeoExecution) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *GeoExecution) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (g GeoExecution) String() string { _ = "STUB: not implemented"; return "" }
