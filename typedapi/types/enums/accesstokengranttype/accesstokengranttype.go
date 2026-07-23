package accesstokengranttype

type AccessTokenGrantType struct {
	Name string
}

var (
	Password = AccessTokenGrantType{"password"}

	Clientcredentials = AccessTokenGrantType{"client_credentials"}

	Kerberos = AccessTokenGrantType{"_kerberos"}

	Refreshtoken = AccessTokenGrantType{"refresh_token"}
)

func (a AccessTokenGrantType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AccessTokenGrantType) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (a AccessTokenGrantType) String() string { _ = "STUB: not implemented"; return "" }
