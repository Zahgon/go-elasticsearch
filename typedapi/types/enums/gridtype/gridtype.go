package gridtype

type GridType struct {
	Name string
}

var (
	Grid = GridType{"grid"}

	Point = GridType{"point"}

	Centroid = GridType{"centroid"}
)

func (g GridType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *GridType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (g GridType) String() string { _ = "STUB: not implemented"; return "" }
