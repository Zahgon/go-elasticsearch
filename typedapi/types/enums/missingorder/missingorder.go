package missingorder

type MissingOrder struct {
	Name string
}

var (
	First = MissingOrder{"first"}

	Last = MissingOrder{"last"}

	Default = MissingOrder{"default"}
)

func (m MissingOrder) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *MissingOrder) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (m MissingOrder) String() string { _ = "STUB: not implemented"; return "" }
