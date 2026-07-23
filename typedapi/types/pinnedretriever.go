package types

type PinnedRetriever struct {
	Docs []SpecifiedDocument `json:"docs,omitempty"`

	Filter []Query  `json:"filter,omitempty"`
	Ids    []string `json:"ids,omitempty"`

	MinScore *float32 `json:"min_score,omitempty"`

	Name_          *string `json:"_name,omitempty"`
	RankWindowSize *int    `json:"rank_window_size,omitempty"`

	Retriever RetrieverContainer `json:"retriever"`
}

func (s *PinnedRetriever) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewPinnedRetriever() *PinnedRetriever { _ = "STUB: not implemented"; return nil }

type PinnedRetrieverVariant interface {
	PinnedRetrieverCaster() *PinnedRetriever
}

func (s *PinnedRetriever) PinnedRetrieverCaster() *PinnedRetriever {
	_ = "STUB: not implemented"
	return nil
}
