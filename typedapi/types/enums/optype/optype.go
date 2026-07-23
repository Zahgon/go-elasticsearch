package optype

type OpType struct {
	Name string
}

var (
	Index = OpType{"index"}

	Create = OpType{"create"}
)

func (o OpType) MarshalText() (text []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

func (o *OpType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (o OpType) String() string { _ = "STUB: not implemented"; return "" }
