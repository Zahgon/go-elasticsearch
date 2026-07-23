package deepseekservicetype

type DeepSeekServiceType struct {
	Name string
}

var (
	Deepseek = DeepSeekServiceType{"deepseek"}
)

func (d DeepSeekServiceType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DeepSeekServiceType) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (d DeepSeekServiceType) String() string { _ = "STUB: not implemented"; return "" }
