package conflicts

type Conflicts struct {
	Name string
}

var (
	Abort = Conflicts{"abort"}

	Proceed = Conflicts{"proceed"}
)

func (c Conflicts) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Conflicts) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (c Conflicts) String() string { _ = "STUB: not implemented"; return "" }
