package shutdownstatus

type ShutdownStatus struct {
	Name string
}

var (
	Notstarted = ShutdownStatus{"not_started"}

	Inprogress = ShutdownStatus{"in_progress"}

	Stalled = ShutdownStatus{"stalled"}

	Complete = ShutdownStatus{"complete"}
)

func (s ShutdownStatus) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *ShutdownStatus) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (s ShutdownStatus) String() string { _ = "STUB: not implemented"; return "" }
