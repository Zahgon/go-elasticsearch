package sourcemode

type SourceMode struct {
	Name string
}

var (
	Disabled = SourceMode{"disabled"}

	Stored = SourceMode{"stored"}

	Synthetic = SourceMode{"synthetic"}
)

func (s SourceMode) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *SourceMode) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (s SourceMode) String() string { _ = "STUB: not implemented"; return "" }
