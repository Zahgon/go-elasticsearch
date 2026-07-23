package shardroutingstate

type ShardRoutingState struct {
	Name string
}

var (
	UNASSIGNED = ShardRoutingState{"UNASSIGNED"}

	INITIALIZING = ShardRoutingState{"INITIALIZING"}

	STARTED = ShardRoutingState{"STARTED"}

	RELOCATING = ShardRoutingState{"RELOCATING"}
)

func (s ShardRoutingState) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *ShardRoutingState) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (s ShardRoutingState) String() string { _ = "STUB: not implemented"; return "" }
