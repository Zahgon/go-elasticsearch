package apikeytype

type ApiKeyType struct {
	Name string
}

var (
	Rest = ApiKeyType{"rest"}

	Crosscluster = ApiKeyType{"cross_cluster"}
)

func (a ApiKeyType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *ApiKeyType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (a ApiKeyType) String() string { _ = "STUB: not implemented"; return "" }
