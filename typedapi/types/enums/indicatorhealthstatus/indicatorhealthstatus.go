package indicatorhealthstatus

type IndicatorHealthStatus struct {
	Name string
}

var (
	Green = IndicatorHealthStatus{"green"}

	Yellow = IndicatorHealthStatus{"yellow"}

	Red = IndicatorHealthStatus{"red"}

	Unknown = IndicatorHealthStatus{"unknown"}

	Unavailable = IndicatorHealthStatus{"unavailable"}
)

func (i IndicatorHealthStatus) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *IndicatorHealthStatus) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (i IndicatorHealthStatus) String() string { _ = "STUB: not implemented"; return "" }
