package tasktypeamazonbedrock

type TaskTypeAmazonBedrock struct {
	Name string
}

var (
	Chatcompletion = TaskTypeAmazonBedrock{"chat_completion"}

	Completion = TaskTypeAmazonBedrock{"completion"}

	Textembedding = TaskTypeAmazonBedrock{"text_embedding"}
)

func (t TaskTypeAmazonBedrock) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TaskTypeAmazonBedrock) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (t TaskTypeAmazonBedrock) String() string { _ = "STUB: not implemented"; return "" }
