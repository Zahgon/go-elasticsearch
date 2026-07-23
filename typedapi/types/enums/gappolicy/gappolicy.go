package gappolicy

type GapPolicy struct {
	Name string
}

var (
	Skip = GapPolicy{"skip"}

	Insertzeros = GapPolicy{"insert_zeros"}

	Keepvalues = GapPolicy{"keep_values"}
)

func (g GapPolicy) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *GapPolicy) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (g GapPolicy) String() string { _ = "STUB: not implemented"; return "" }
