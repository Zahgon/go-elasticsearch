package syntheticsourcekeepenum

type SyntheticSourceKeepEnum struct {
	Name string
}

var (
	None = SyntheticSourceKeepEnum{"none"}

	Arrays = SyntheticSourceKeepEnum{"arrays"}

	All = SyntheticSourceKeepEnum{"all"}
)

func (s SyntheticSourceKeepEnum) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *SyntheticSourceKeepEnum) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s SyntheticSourceKeepEnum) String() string { _ = "STUB: not implemented"; return "" }
