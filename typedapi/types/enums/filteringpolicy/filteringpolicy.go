package filteringpolicy

type FilteringPolicy struct {
	Name string
}

var (
	Exclude = FilteringPolicy{"exclude"}

	Include = FilteringPolicy{"include"}
)

func (f FilteringPolicy) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FilteringPolicy) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (f FilteringPolicy) String() string { _ = "STUB: not implemented"; return "" }
