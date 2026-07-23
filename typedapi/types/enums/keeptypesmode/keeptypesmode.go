package keeptypesmode

type KeepTypesMode struct {
	Name string
}

var (
	Include = KeepTypesMode{"include"}

	Exclude = KeepTypesMode{"exclude"}
)

func (k KeepTypesMode) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *KeepTypesMode) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (k KeepTypesMode) String() string { _ = "STUB: not implemented"; return "" }
