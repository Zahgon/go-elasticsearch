package apikeygranttype

type ApiKeyGrantType struct {
	Name string
}

var (
	Accesstoken = ApiKeyGrantType{"access_token"}

	Password = ApiKeyGrantType{"password"}
)

func (a ApiKeyGrantType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *ApiKeyGrantType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (a ApiKeyGrantType) String() string { _ = "STUB: not implemented"; return "" }
