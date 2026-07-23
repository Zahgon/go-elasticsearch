package jinaaielementtype

type JinaAIElementType struct {
	Name string
}

var (
	Binary = JinaAIElementType{"binary"}

	Bit = JinaAIElementType{"bit"}

	Float = JinaAIElementType{"float"}
)

func (j JinaAIElementType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (j *JinaAIElementType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (j JinaAIElementType) String() string { _ = "STUB: not implemented"; return "" }
