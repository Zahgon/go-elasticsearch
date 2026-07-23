package holtwinterstype

type HoltWintersType struct {
	Name string
}

var (
	Additive = HoltWintersType{"add"}

	Multiplicative = HoltWintersType{"mult"}
)

func (h HoltWintersType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *HoltWintersType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (h HoltWintersType) String() string { _ = "STUB: not implemented"; return "" }
