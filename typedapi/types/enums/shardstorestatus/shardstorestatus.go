package shardstorestatus

type ShardStoreStatus struct {
	Name string
}

var (
	Green = ShardStoreStatus{"green"}

	Yellow = ShardStoreStatus{"yellow"}

	Red = ShardStoreStatus{"red"}

	All = ShardStoreStatus{"all"}
)

func (s ShardStoreStatus) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *ShardStoreStatus) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (s ShardStoreStatus) String() string { _ = "STUB: not implemented"; return "" }
