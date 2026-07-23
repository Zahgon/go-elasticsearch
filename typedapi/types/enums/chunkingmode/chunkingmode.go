package chunkingmode

type ChunkingMode struct {
	Name string
}

var (
	Auto = ChunkingMode{"auto"}

	Manual = ChunkingMode{"manual"}

	Off = ChunkingMode{"off"}
)

func (c ChunkingMode) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ChunkingMode) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (c ChunkingMode) String() string { _ = "STUB: not implemented"; return "" }
