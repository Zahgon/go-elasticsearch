package tasktypeamazonsagemaker

type TaskTypeAmazonSageMaker struct {
	Name string
}

var (
	Textembedding = TaskTypeAmazonSageMaker{"text_embedding"}

	Completion = TaskTypeAmazonSageMaker{"completion"}

	Chatcompletion = TaskTypeAmazonSageMaker{"chat_completion"}

	Sparseembedding = TaskTypeAmazonSageMaker{"sparse_embedding"}

	Rerank = TaskTypeAmazonSageMaker{"rerank"}
)

func (t TaskTypeAmazonSageMaker) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TaskTypeAmazonSageMaker) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (t TaskTypeAmazonSageMaker) String() string { _ = "STUB: not implemented"; return "" }
