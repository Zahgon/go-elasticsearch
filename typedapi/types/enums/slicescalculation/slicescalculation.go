package slicescalculation

type SlicesCalculation struct {
	Name string
}

var (
	Auto = SlicesCalculation{"auto"}
)

func (s SlicesCalculation) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *SlicesCalculation) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (s SlicesCalculation) String() string { _ = "STUB: not implemented"; return "" }
