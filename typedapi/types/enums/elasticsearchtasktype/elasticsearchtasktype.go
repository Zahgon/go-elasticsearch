package elasticsearchtasktype

type ElasticsearchTaskType struct {
	Name string
}

var (
	Rerank = ElasticsearchTaskType{"rerank"}

	Sparseembedding = ElasticsearchTaskType{"sparse_embedding"}

	Textembedding = ElasticsearchTaskType{"text_embedding"}
)

func (e ElasticsearchTaskType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *ElasticsearchTaskType) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (e ElasticsearchTaskType) String() string { _ = "STUB: not implemented"; return "" }
