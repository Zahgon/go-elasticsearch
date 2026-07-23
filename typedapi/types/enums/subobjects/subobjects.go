package subobjects

type Subobjects struct {
	Name string
}

var (
	True = Subobjects{"true"}

	False = Subobjects{"false"}

	Auto = Subobjects{"auto"}
)

func (s *Subobjects) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s Subobjects) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Subobjects) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (s Subobjects) String() string { _ = "STUB: not implemented"; return "" }
