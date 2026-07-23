package boundaryscanner

type BoundaryScanner struct {
	Name string
}

var (
	Chars = BoundaryScanner{"chars"}

	Sentence = BoundaryScanner{"sentence"}

	Word = BoundaryScanner{"word"}
)

func (b BoundaryScanner) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *BoundaryScanner) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (b BoundaryScanner) String() string { _ = "STUB: not implemented"; return "" }
