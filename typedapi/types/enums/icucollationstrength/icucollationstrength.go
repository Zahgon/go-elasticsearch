package icucollationstrength

type IcuCollationStrength struct {
	Name string
}

var (
	Primary = IcuCollationStrength{"primary"}

	Secondary = IcuCollationStrength{"secondary"}

	Tertiary = IcuCollationStrength{"tertiary"}

	Quaternary = IcuCollationStrength{"quaternary"}

	Identical = IcuCollationStrength{"identical"}
)

func (i IcuCollationStrength) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *IcuCollationStrength) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (i IcuCollationStrength) String() string { _ = "STUB: not implemented"; return "" }
