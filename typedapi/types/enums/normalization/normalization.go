package normalization

type Normalization struct {
	Name string
}

var (
	No = Normalization{"no"}

	H1 = Normalization{"h1"}

	H2 = Normalization{"h2"}

	H3 = Normalization{"h3"}

	Z = Normalization{"z"}
)

func (n Normalization) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (n *Normalization) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (n Normalization) String() string { _ = "STUB: not implemented"; return "" }
