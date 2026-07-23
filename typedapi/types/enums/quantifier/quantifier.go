package quantifier

type Quantifier struct {
	Name string
}

var (
	Some = Quantifier{"some"}

	All = Quantifier{"all"}
)

func (q Quantifier) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (q *Quantifier) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (q Quantifier) String() string { _ = "STUB: not implemented"; return "" }
