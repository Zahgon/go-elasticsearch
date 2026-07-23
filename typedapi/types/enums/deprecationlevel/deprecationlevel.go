package deprecationlevel

type DeprecationLevel struct {
	Name string
}

var (
	None = DeprecationLevel{"none"}

	Info = DeprecationLevel{"info"}

	Warning = DeprecationLevel{"warning"}

	Critical = DeprecationLevel{"critical"}
)

func (d DeprecationLevel) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DeprecationLevel) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (d DeprecationLevel) String() string { _ = "STUB: not implemented"; return "" }
