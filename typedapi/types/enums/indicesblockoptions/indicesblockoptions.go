package indicesblockoptions

type IndicesBlockOptions struct {
	Name string
}

var (
	Metadata = IndicesBlockOptions{"metadata"}

	Read = IndicesBlockOptions{"read"}

	Readonly = IndicesBlockOptions{"read_only"}

	Write = IndicesBlockOptions{"write"}
)

func (i IndicesBlockOptions) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *IndicesBlockOptions) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (i IndicesBlockOptions) String() string { _ = "STUB: not implemented"; return "" }
