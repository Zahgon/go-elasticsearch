package reasoningsummary

type ReasoningSummary struct {
	Name string
}

var (
	Auto = ReasoningSummary{"auto"}

	Concise = ReasoningSummary{"concise"}

	Detailed = ReasoningSummary{"detailed"}
)

func (r ReasoningSummary) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *ReasoningSummary) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (r ReasoningSummary) String() string { _ = "STUB: not implemented"; return "" }
