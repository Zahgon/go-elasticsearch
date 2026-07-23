package dfrbasicmodel

type DFRBasicModel struct {
	Name string
}

var (
	Be = DFRBasicModel{"be"}

	D = DFRBasicModel{"d"}

	G = DFRBasicModel{"g"}

	If = DFRBasicModel{"if"}

	In = DFRBasicModel{"in"}

	Ine = DFRBasicModel{"ine"}

	P = DFRBasicModel{"p"}
)

func (d DFRBasicModel) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DFRBasicModel) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (d DFRBasicModel) String() string { _ = "STUB: not implemented"; return "" }
