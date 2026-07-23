package highlighterencoder

type HighlighterEncoder struct {
	Name string
}

var (
	Default = HighlighterEncoder{"default"}

	Html = HighlighterEncoder{"html"}
)

func (h HighlighterEncoder) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *HighlighterEncoder) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (h HighlighterEncoder) String() string { _ = "STUB: not implemented"; return "" }
