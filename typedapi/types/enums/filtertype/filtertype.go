package filtertype

type FilterType struct {
	Name string
}

var (
	Include = FilterType{"include"}

	Exclude = FilterType{"exclude"}
)

func (f FilterType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FilterType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (f FilterType) String() string { _ = "STUB: not implemented"; return "" }
