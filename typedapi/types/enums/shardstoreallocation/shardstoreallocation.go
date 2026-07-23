package shardstoreallocation

type ShardStoreAllocation struct {
	Name string
}

var (
	Primary = ShardStoreAllocation{"primary"}

	Replica = ShardStoreAllocation{"replica"}

	Unused = ShardStoreAllocation{"unused"}
)

func (s ShardStoreAllocation) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *ShardStoreAllocation) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s ShardStoreAllocation) String() string { _ = "STUB: not implemented"; return "" }
