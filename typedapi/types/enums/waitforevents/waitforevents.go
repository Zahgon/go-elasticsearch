package waitforevents

type WaitForEvents struct {
	Name string
}

var (
	Immediate = WaitForEvents{"immediate"}

	Urgent = WaitForEvents{"urgent"}

	High = WaitForEvents{"high"}

	Normal = WaitForEvents{"normal"}

	Low = WaitForEvents{"low"}

	Languid = WaitForEvents{"languid"}
)

func (w WaitForEvents) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *WaitForEvents) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (w WaitForEvents) String() string { _ = "STUB: not implemented"; return "" }
