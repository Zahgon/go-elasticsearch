package icunormalizationmode

type IcuNormalizationMode struct {
	Name string
}

var (
	Decompose = IcuNormalizationMode{"decompose"}

	Compose = IcuNormalizationMode{"compose"}
)

func (i IcuNormalizationMode) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *IcuNormalizationMode) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (i IcuNormalizationMode) String() string { _ = "STUB: not implemented"; return "" }
