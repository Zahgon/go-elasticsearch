package llamaservicetype

type LlamaServiceType struct {
	Name string
}

var (
	Llama = LlamaServiceType{"llama"}
)

func (l LlamaServiceType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *LlamaServiceType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (l LlamaServiceType) String() string { _ = "STUB: not implemented"; return "" }
