package lowercasetokenfilterlanguages

type LowercaseTokenFilterLanguages struct {
	Name string
}

var (
	Greek = LowercaseTokenFilterLanguages{"greek"}

	Irish = LowercaseTokenFilterLanguages{"irish"}

	Turkish = LowercaseTokenFilterLanguages{"turkish"}
)

func (l LowercaseTokenFilterLanguages) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *LowercaseTokenFilterLanguages) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (l LowercaseTokenFilterLanguages) String() string { _ = "STUB: not implemented"; return "" }
