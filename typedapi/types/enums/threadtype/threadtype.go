package threadtype

type ThreadType struct {
	Name string
}

var (
	Cpu = ThreadType{"cpu"}

	Wait = ThreadType{"wait"}

	Block = ThreadType{"block"}

	Gpu = ThreadType{"gpu"}

	Mem = ThreadType{"mem"}
)

func (t ThreadType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *ThreadType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (t ThreadType) String() string { _ = "STUB: not implemented"; return "" }
