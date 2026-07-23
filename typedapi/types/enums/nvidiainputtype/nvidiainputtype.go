package nvidiainputtype

type NvidiaInputType struct {
	Name string
}

var (
	Ingest = NvidiaInputType{"ingest"}

	Search = NvidiaInputType{"search"}
)

func (n NvidiaInputType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (n *NvidiaInputType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (n NvidiaInputType) String() string { _ = "STUB: not implemented"; return "" }
