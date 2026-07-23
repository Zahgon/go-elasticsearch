package googleaistudiotasktype

type GoogleAiStudioTaskType struct {
	Name string
}

var (
	Completion = GoogleAiStudioTaskType{"completion"}

	Textembedding = GoogleAiStudioTaskType{"text_embedding"}
)

func (g GoogleAiStudioTaskType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *GoogleAiStudioTaskType) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (g GoogleAiStudioTaskType) String() string { _ = "STUB: not implemented"; return "" }
