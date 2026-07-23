package fieldaccesspattern

type FieldAccessPattern struct {
	Name string
}

var (
	Classic = FieldAccessPattern{"classic"}

	Flexible = FieldAccessPattern{"flexible"}
)

func (f FieldAccessPattern) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FieldAccessPattern) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (f FieldAccessPattern) String() string { _ = "STUB: not implemented"; return "" }
