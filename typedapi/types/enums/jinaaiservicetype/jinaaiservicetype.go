package jinaaiservicetype

type JinaAIServiceType struct {
	Name string
}

var (
	Jinaai = JinaAIServiceType{"jinaai"}
)

func (j JinaAIServiceType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (j *JinaAIServiceType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (j JinaAIServiceType) String() string { _ = "STUB: not implemented"; return "" }
