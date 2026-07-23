package dfraftereffect

type DFRAfterEffect struct {
	Name string
}

var (
	No = DFRAfterEffect{"no"}

	B = DFRAfterEffect{"b"}

	L = DFRAfterEffect{"l"}
)

func (d DFRAfterEffect) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DFRAfterEffect) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (d DFRAfterEffect) String() string { _ = "STUB: not implemented"; return "" }
