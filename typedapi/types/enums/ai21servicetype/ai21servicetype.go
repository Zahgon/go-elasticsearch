package ai21servicetype

type Ai21ServiceType struct {
	Name string
}

var (
	Ai21 = Ai21ServiceType{"ai21"}
)

func (a Ai21ServiceType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *Ai21ServiceType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (a Ai21ServiceType) String() string { _ = "STUB: not implemented"; return "" }
