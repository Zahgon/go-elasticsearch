package kuromojitokenizationmode

type KuromojiTokenizationMode struct {
	Name string
}

var (
	Normal = KuromojiTokenizationMode{"normal"}

	Search = KuromojiTokenizationMode{"search"}

	Extended = KuromojiTokenizationMode{"extended"}
)

func (k KuromojiTokenizationMode) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *KuromojiTokenizationMode) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (k KuromojiTokenizationMode) String() string { _ = "STUB: not implemented"; return "" }
