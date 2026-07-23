package operationtype

type OperationType struct {
	Name string
}

var (
	Index = OperationType{"index"}

	Create = OperationType{"create"}

	Update = OperationType{"update"}

	Delete = OperationType{"delete"}
)

func (o OperationType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o *OperationType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (o OperationType) String() string { _ = "STUB: not implemented"; return "" }
