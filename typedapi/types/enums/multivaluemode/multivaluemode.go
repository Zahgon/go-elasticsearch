package multivaluemode

type MultiValueMode struct {
	Name string
}

var (
	Min = MultiValueMode{"min"}

	Max = MultiValueMode{"max"}

	Avg = MultiValueMode{"avg"}

	Sum = MultiValueMode{"sum"}
)

func (m MultiValueMode) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *MultiValueMode) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (m MultiValueMode) String() string { _ = "STUB: not implemented"; return "" }
