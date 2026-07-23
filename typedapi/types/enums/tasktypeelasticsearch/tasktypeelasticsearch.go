package tasktypeelasticsearch

type TaskTypeElasticsearch struct {
	Name string
}

var (
	Sparseembedding = TaskTypeElasticsearch{"sparse_embedding"}

	Textembedding = TaskTypeElasticsearch{"text_embedding"}

	Rerank = TaskTypeElasticsearch{"rerank"}
)

func (t TaskTypeElasticsearch) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TaskTypeElasticsearch) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (t TaskTypeElasticsearch) String() string { _ = "STUB: not implemented"; return "" }
