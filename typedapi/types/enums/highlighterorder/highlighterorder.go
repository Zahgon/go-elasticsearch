package highlighterorder

type HighlighterOrder struct {
	Name string
}

var (
	Score = HighlighterOrder{"score"}
)

func (h HighlighterOrder) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *HighlighterOrder) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (h HighlighterOrder) String() string { _ = "STUB: not implemented"; return "" }
