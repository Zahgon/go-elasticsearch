package huggingfaceservicetype

type HuggingFaceServiceType struct {
	Name string
}

var (
	Huggingface = HuggingFaceServiceType{"hugging_face"}
)

func (h HuggingFaceServiceType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *HuggingFaceServiceType) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (h HuggingFaceServiceType) String() string { _ = "STUB: not implemented"; return "" }
