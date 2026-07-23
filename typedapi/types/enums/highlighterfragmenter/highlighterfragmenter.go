package highlighterfragmenter

type HighlighterFragmenter struct {
	Name string
}

var (
	Simple = HighlighterFragmenter{"simple"}

	Span = HighlighterFragmenter{"span"}
)

func (h HighlighterFragmenter) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *HighlighterFragmenter) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (h HighlighterFragmenter) String() string { _ = "STUB: not implemented"; return "" }
