package edgengramside

type EdgeNGramSide struct {
	Name string
}

var (
	Front = EdgeNGramSide{"front"}

	Back = EdgeNGramSide{"back"}
)

func (e EdgeNGramSide) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *EdgeNGramSide) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (e EdgeNGramSide) String() string { _ = "STUB: not implemented"; return "" }
