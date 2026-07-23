package streamtype

type StreamType struct {
	Name string
}

var (
	Logs = StreamType{"logs"}

	Logsotel = StreamType{"logs.otel"}

	Logsecs = StreamType{"logs.ecs"}
)

func (s StreamType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *StreamType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (s StreamType) String() string { _ = "STUB: not implemented"; return "" }
