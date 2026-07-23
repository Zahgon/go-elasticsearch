package granttype

type GrantType struct {
	Name string
}

var (
	Password = GrantType{"password"}

	Accesstoken = GrantType{"access_token"}
)

func (g GrantType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *GrantType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (g GrantType) String() string { _ = "STUB: not implemented"; return "" }
