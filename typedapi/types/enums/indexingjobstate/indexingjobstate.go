package indexingjobstate

type IndexingJobState struct {
	Name string
}

var (
	Started = IndexingJobState{"started"}

	Indexing = IndexingJobState{"indexing"}

	Stopping = IndexingJobState{"stopping"}

	Stopped = IndexingJobState{"stopped"}

	Aborting = IndexingJobState{"aborting"}
)

func (i IndexingJobState) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *IndexingJobState) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (i IndexingJobState) String() string { _ = "STUB: not implemented"; return "" }
