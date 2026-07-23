package icucollationdecomposition

type IcuCollationDecomposition struct {
	Name string
}

var (
	No = IcuCollationDecomposition{"no"}

	Identical = IcuCollationDecomposition{"identical"}
)

func (i IcuCollationDecomposition) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *IcuCollationDecomposition) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (i IcuCollationDecomposition) String() string { _ = "STUB: not implemented"; return "" }
