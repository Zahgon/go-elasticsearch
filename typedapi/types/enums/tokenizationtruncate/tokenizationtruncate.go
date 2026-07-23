package tokenizationtruncate

type TokenizationTruncate struct {
	Name string
}

var (
	First = TokenizationTruncate{"first"}

	Second = TokenizationTruncate{"second"}

	None = TokenizationTruncate{"none"}
)

func (t TokenizationTruncate) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TokenizationTruncate) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (t TokenizationTruncate) String() string { _ = "STUB: not implemented"; return "" }
