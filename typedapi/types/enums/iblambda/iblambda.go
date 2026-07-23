package iblambda

type IBLambda struct {
	Name string
}

var (
	Df = IBLambda{"df"}

	Ttf = IBLambda{"ttf"}
)

func (i IBLambda) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *IBLambda) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (i IBLambda) String() string { _ = "STUB: not implemented"; return "" }
