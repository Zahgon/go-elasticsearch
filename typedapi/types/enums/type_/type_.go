package type_

type Type struct {
	Name string
}

var (
	Restart = Type{"restart"}

	Remove = Type{"remove"}

	Replace = Type{"replace"}
)

func (t Type) MarshalText() (text []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

func (t *Type) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (t Type) String() string { _ = "STUB: not implemented"; return "" }
