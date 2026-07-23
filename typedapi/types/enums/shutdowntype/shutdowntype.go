package shutdowntype

type ShutdownType struct {
	Name string
}

var (
	Remove = ShutdownType{"remove"}

	Restart = ShutdownType{"restart"}
)

func (s ShutdownType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *ShutdownType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (s ShutdownType) String() string { _ = "STUB: not implemented"; return "" }
