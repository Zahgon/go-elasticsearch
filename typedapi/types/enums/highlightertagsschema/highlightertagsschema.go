package highlightertagsschema

type HighlighterTagsSchema struct {
	Name string
}

var (
	Styled = HighlighterTagsSchema{"styled"}
)

func (h HighlighterTagsSchema) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *HighlighterTagsSchema) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (h HighlighterTagsSchema) String() string { _ = "STUB: not implemented"; return "" }
