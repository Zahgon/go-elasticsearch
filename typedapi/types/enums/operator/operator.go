package operator

type Operator struct {
	Name string
}

var (
	And = Operator{"and"}

	Or = Operator{"or"}
)

func (o Operator) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o *Operator) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (o Operator) String() string { _ = "STUB: not implemented"; return "" }
