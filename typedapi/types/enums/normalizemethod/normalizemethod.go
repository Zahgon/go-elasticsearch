package normalizemethod

type NormalizeMethod struct {
	Name string
}

var (
	Rescale01 = NormalizeMethod{"rescale_0_1"}

	Rescale0100 = NormalizeMethod{"rescale_0_100"}

	Percentofsum = NormalizeMethod{"percent_of_sum"}

	Mean = NormalizeMethod{"mean"}

	Zscore = NormalizeMethod{"z-score"}

	Softmax = NormalizeMethod{"softmax"}
)

func (n NormalizeMethod) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (n *NormalizeMethod) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (n NormalizeMethod) String() string { _ = "STUB: not implemented"; return "" }
