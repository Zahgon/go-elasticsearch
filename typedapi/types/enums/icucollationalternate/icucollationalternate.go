package icucollationalternate

type IcuCollationAlternate struct {
	Name string
}

var (
	Shifted = IcuCollationAlternate{"shifted"}

	NonIgnorable = IcuCollationAlternate{"non-ignorable"}
)

func (i IcuCollationAlternate) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *IcuCollationAlternate) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (i IcuCollationAlternate) String() string { _ = "STUB: not implemented"; return "" }
