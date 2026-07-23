package excludefrequent

type ExcludeFrequent struct {
	Name string
}

var (
	All = ExcludeFrequent{"all"}

	None = ExcludeFrequent{"none"}

	By = ExcludeFrequent{"by"}

	Over = ExcludeFrequent{"over"}
)

func (e ExcludeFrequent) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *ExcludeFrequent) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (e ExcludeFrequent) String() string { _ = "STUB: not implemented"; return "" }
