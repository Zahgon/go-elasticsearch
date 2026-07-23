package modeenum

type ModeEnum struct {
	Name string
}

var (
	Upgrade = ModeEnum{"upgrade"}
)

func (m ModeEnum) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *ModeEnum) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (m ModeEnum) String() string { _ = "STUB: not implemented"; return "" }
