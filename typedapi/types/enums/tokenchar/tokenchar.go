package tokenchar

type TokenChar struct {
	Name string
}

var (
	Letter = TokenChar{"letter"}

	Digit = TokenChar{"digit"}

	Whitespace = TokenChar{"whitespace"}

	Punctuation = TokenChar{"punctuation"}

	Symbol = TokenChar{"symbol"}

	Custom = TokenChar{"custom"}
)

func (t TokenChar) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TokenChar) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (t TokenChar) String() string { _ = "STUB: not implemented"; return "" }
