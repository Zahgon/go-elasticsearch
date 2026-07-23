package geogridtiletype

type GeoGridTileType struct {
	Name string
}

var (
	Geotile = GeoGridTileType{"geotile"}

	Geohex = GeoGridTileType{"geohex"}

	Geohash = GeoGridTileType{"geohash"}
)

func (g GeoGridTileType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *GeoGridTileType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (g GeoGridTileType) String() string { _ = "STUB: not implemented"; return "" }
