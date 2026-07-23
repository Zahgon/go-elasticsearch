package statslevel

type StatsLevel struct {
	Name string
}

var (
	Cluster = StatsLevel{"cluster"}

	Indices = StatsLevel{"indices"}

	Shards = StatsLevel{"shards"}
)

func (s StatsLevel) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *StatsLevel) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (s StatsLevel) String() string { _ = "STUB: not implemented"; return "" }
