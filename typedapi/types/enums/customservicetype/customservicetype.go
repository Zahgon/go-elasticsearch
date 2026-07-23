package customservicetype

type CustomServiceType struct {
	Name string
}

var (
	Custom = CustomServiceType{"custom"}
)

func (c CustomServiceType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CustomServiceType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (c CustomServiceType) String() string { _ = "STUB: not implemented"; return "" }
