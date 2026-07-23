package tasktypealibabacloudai

type TaskTypeAlibabaCloudAI struct {
	Name string
}

var (
	Textembedding = TaskTypeAlibabaCloudAI{"text_embedding"}

	Rerank = TaskTypeAlibabaCloudAI{"rerank"}

	Completion = TaskTypeAlibabaCloudAI{"completion"}

	Sparseembedding = TaskTypeAlibabaCloudAI{"sparse_embedding"}
)

func (t TaskTypeAlibabaCloudAI) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TaskTypeAlibabaCloudAI) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (t TaskTypeAlibabaCloudAI) String() string { _ = "STUB: not implemented"; return "" }
