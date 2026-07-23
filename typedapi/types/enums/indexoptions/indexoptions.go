package indexoptions

type IndexOptions struct {
	Name string
}

var (
	Docs = IndexOptions{"docs"}

	Freqs = IndexOptions{"freqs"}

	Positions = IndexOptions{"positions"}

	Offsets = IndexOptions{"offsets"}
)

func (i IndexOptions) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *IndexOptions) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (i IndexOptions) String() string { _ = "STUB: not implemented"; return "" }
