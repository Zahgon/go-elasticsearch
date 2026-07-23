package shardstate

type ShardState struct {
	Name string
}

var (
	INIT = ShardState{"INIT"}

	SUCCESS = ShardState{"SUCCESS"}

	FAILED = ShardState{"FAILED"}

	ABORTED = ShardState{"ABORTED"}

	MISSING = ShardState{"MISSING"}

	WAITING = ShardState{"WAITING"}

	QUEUED = ShardState{"QUEUED"}

	PAUSEDFORNODEREMOVAL = ShardState{"PAUSED_FOR_NODE_REMOVAL"}
)

func (s ShardState) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *ShardState) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (s ShardState) String() string { _ = "STUB: not implemented"; return "" }
