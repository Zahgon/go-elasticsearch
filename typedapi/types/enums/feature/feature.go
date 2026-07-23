package feature

type Feature struct {
	Name string
}

var (
	Aliases = Feature{"aliases"}

	Mappings = Feature{"mappings"}

	Settings = Feature{"settings"}
)

func (f Feature) MarshalText() (text []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

func (f *Feature) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (f Feature) String() string { _ = "STUB: not implemented"; return "" }
