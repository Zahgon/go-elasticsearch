package level

type Level struct {
	Name string
}

var (
	Cluster = Level{"cluster"}

	Indices = Level{"indices"}

	Shards = Level{"shards"}
)

func (l Level) MarshalText() (text []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

func (l *Level) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (l Level) String() string { _ = "STUB: not implemented"; return "" }
