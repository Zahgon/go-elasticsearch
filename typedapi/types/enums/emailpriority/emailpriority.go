package emailpriority

type EmailPriority struct {
	Name string
}

var (
	Lowest = EmailPriority{"lowest"}

	Low = EmailPriority{"low"}

	Normal = EmailPriority{"normal"}

	High = EmailPriority{"high"}

	Highest = EmailPriority{"highest"}
)

func (e EmailPriority) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *EmailPriority) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (e EmailPriority) String() string { _ = "STUB: not implemented"; return "" }
