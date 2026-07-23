package gridaggregationtype

type GridAggregationType struct {
	Name string
}

var (
	Geotile = GridAggregationType{"geotile"}

	Geohex = GridAggregationType{"geohex"}
)

func (g GridAggregationType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *GridAggregationType) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (g GridAggregationType) String() string { _ = "STUB: not implemented"; return "" }
