package alibabacloudtasktype

type AlibabaCloudTaskType struct {
	Name string
}

var (
	Completion = AlibabaCloudTaskType{"completion"}

	Rerank = AlibabaCloudTaskType{"rerank"}

	Sparseembedding = AlibabaCloudTaskType{"sparse_embedding"}

	Textembedding = AlibabaCloudTaskType{"text_embedding"}
)

func (a AlibabaCloudTaskType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AlibabaCloudTaskType) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (a AlibabaCloudTaskType) String() string { _ = "STUB: not implemented"; return "" }
