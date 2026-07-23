package nvidiaservicetype

type NvidiaServiceType struct {
	Name string
}

var (
	Nvidia = NvidiaServiceType{"nvidia"}
)

func (n NvidiaServiceType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (n *NvidiaServiceType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (n NvidiaServiceType) String() string { _ = "STUB: not implemented"; return "" }
