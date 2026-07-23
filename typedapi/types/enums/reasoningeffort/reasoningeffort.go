package reasoningeffort

type ReasoningEffort struct {
	Name string
}

var (
	Xhigh = ReasoningEffort{"xhigh"}

	High = ReasoningEffort{"high"}

	Medium = ReasoningEffort{"medium"}

	Low = ReasoningEffort{"low"}

	Minimal = ReasoningEffort{"minimal"}

	None = ReasoningEffort{"none"}
)

func (r ReasoningEffort) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *ReasoningEffort) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (r ReasoningEffort) String() string { _ = "STUB: not implemented"; return "" }
