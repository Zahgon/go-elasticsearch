package ratemode

type RateMode struct {
	Name string
}

var (
	Sum = RateMode{"sum"}

	Valuecount = RateMode{"value_count"}
)

func (r RateMode) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *RateMode) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (r RateMode) String() string { _ = "STUB: not implemented"; return "" }
