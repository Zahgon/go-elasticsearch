package healthstatus

type HealthStatus struct {
	Name string
}

var (
	Green = HealthStatus{"green"}

	Yellow = HealthStatus{"yellow"}

	Red = HealthStatus{"red"}

	Unknown = HealthStatus{"unknown"}

	Unavailable = HealthStatus{"unavailable"}
)

func (h HealthStatus) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *HealthStatus) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (h HealthStatus) String() string { _ = "STUB: not implemented"; return "" }
