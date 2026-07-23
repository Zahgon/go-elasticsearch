package cjkbigramignoredscript

type CjkBigramIgnoredScript struct {
	Name string
}

var (
	Han = CjkBigramIgnoredScript{"han"}

	Hangul = CjkBigramIgnoredScript{"hangul"}

	Hiragana = CjkBigramIgnoredScript{"hiragana"}

	Katakana = CjkBigramIgnoredScript{"katakana"}
)

func (c CjkBigramIgnoredScript) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CjkBigramIgnoredScript) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (c CjkBigramIgnoredScript) String() string { _ = "STUB: not implemented"; return "" }
