package icucollationcasefirst

type IcuCollationCaseFirst struct {
	Name string
}

var (
	Lower = IcuCollationCaseFirst{"lower"}

	Upper = IcuCollationCaseFirst{"upper"}
)

func (i IcuCollationCaseFirst) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *IcuCollationCaseFirst) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (i IcuCollationCaseFirst) String() string { _ = "STUB: not implemented"; return "" }
