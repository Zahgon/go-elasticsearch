package highlightertype

type HighlighterType struct {
	Name string
}

var (
	Plain = HighlighterType{"plain"}

	Fastvector = HighlighterType{"fvh"}

	Unified = HighlighterType{"unified"}
)

func (h HighlighterType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *HighlighterType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (h HighlighterType) String() string { _ = "STUB: not implemented"; return "" }
