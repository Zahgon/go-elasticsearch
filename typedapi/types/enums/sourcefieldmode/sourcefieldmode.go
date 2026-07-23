package sourcefieldmode

type SourceFieldMode struct {
	Name string
}

var (
	Disabled = SourceFieldMode{"disabled"}

	Stored = SourceFieldMode{"stored"}

	Synthetic = SourceFieldMode{"synthetic"}
)

func (s SourceFieldMode) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *SourceFieldMode) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (s SourceFieldMode) String() string { _ = "STUB: not implemented"; return "" }
