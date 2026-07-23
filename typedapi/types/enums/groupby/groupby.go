package groupby

type GroupBy struct {
	Name string
}

var (
	Nodes = GroupBy{"nodes"}

	Parents = GroupBy{"parents"}

	None = GroupBy{"none"}
)

func (g GroupBy) MarshalText() (text []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

func (g *GroupBy) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (g GroupBy) String() string { _ = "STUB: not implemented"; return "" }
