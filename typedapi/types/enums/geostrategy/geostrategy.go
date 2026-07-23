package geostrategy

type GeoStrategy struct {
	Name string
}

var (
	Recursive = GeoStrategy{"recursive"}

	Term = GeoStrategy{"term"}
)

func (g GeoStrategy) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *GeoStrategy) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (g GeoStrategy) String() string { _ = "STUB: not implemented"; return "" }
