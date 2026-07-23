package sortmode

type SortMode struct {
	Name string
}

var (
	Min = SortMode{"min"}

	Max = SortMode{"max"}

	Sum = SortMode{"sum"}

	Avg = SortMode{"avg"}

	Median = SortMode{"median"}
)

func (s SortMode) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *SortMode) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (s SortMode) String() string { _ = "STUB: not implemented"; return "" }
