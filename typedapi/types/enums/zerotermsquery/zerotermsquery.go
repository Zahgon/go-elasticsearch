package zerotermsquery

type ZeroTermsQuery struct {
	Name string
}

var (
	All = ZeroTermsQuery{"all"}

	None = ZeroTermsQuery{"none"}
)

func (z ZeroTermsQuery) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (z *ZeroTermsQuery) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (z ZeroTermsQuery) String() string { _ = "STUB: not implemented"; return "" }
