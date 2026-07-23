package geoshaperelation

type GeoShapeRelation struct {
	Name string
}

var (
	Intersects = GeoShapeRelation{"intersects"}

	Disjoint = GeoShapeRelation{"disjoint"}

	Within = GeoShapeRelation{"within"}

	Contains = GeoShapeRelation{"contains"}
)

func (g GeoShapeRelation) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *GeoShapeRelation) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (g GeoShapeRelation) String() string { _ = "STUB: not implemented"; return "" }
