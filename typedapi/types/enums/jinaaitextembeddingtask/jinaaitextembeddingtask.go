package jinaaitextembeddingtask

type JinaAITextEmbeddingTask struct {
	Name string
}

var (
	Classification = JinaAITextEmbeddingTask{"classification"}

	Clustering = JinaAITextEmbeddingTask{"clustering"}

	Ingest = JinaAITextEmbeddingTask{"ingest"}

	Search = JinaAITextEmbeddingTask{"search"}
)

func (j JinaAITextEmbeddingTask) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (j *JinaAITextEmbeddingTask) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (j JinaAITextEmbeddingTask) String() string { _ = "STUB: not implemented"; return "" }
