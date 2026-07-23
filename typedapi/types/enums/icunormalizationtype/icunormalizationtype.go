package icunormalizationtype

type IcuNormalizationType struct {
	Name string
}

var (
	Nfc = IcuNormalizationType{"nfc"}

	Nfkc = IcuNormalizationType{"nfkc"}

	Nfkccf = IcuNormalizationType{"nfkc_cf"}
)

func (i IcuNormalizationType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *IcuNormalizationType) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (i IcuNormalizationType) String() string { _ = "STUB: not implemented"; return "" }
