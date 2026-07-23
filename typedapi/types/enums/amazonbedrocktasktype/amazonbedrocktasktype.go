package amazonbedrocktasktype

type AmazonBedrockTaskType struct {
	Name string
}

var (
	Chatcompletion = AmazonBedrockTaskType{"chat_completion"}

	Completion = AmazonBedrockTaskType{"completion"}

	Textembedding = AmazonBedrockTaskType{"text_embedding"}
)

func (a AmazonBedrockTaskType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AmazonBedrockTaskType) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (a AmazonBedrockTaskType) String() string { _ = "STUB: not implemented"; return "" }
