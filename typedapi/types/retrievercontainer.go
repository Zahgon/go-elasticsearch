package types

type RetrieverContainer struct {
	Diversify *DiversifyRetriever `json:"diversify,omitempty"`

	Knn *KnnRetriever `json:"knn,omitempty"`

	Linear *LinearRetriever `json:"linear,omitempty"`

	Pinned *PinnedRetriever `json:"pinned,omitempty"`

	Rescorer *RescorerRetriever `json:"rescorer,omitempty"`

	Rrf *RRFRetriever `json:"rrf,omitempty"`

	Rule *RuleRetriever `json:"rule,omitempty"`

	Standard *StandardRetriever `json:"standard,omitempty"`

	TextSimilarityReranker *TextSimilarityReranker `json:"text_similarity_reranker,omitempty"`
}

func NewRetrieverContainer() *RetrieverContainer { _ = "STUB: not implemented"; return nil }

type RetrieverContainerVariant interface {
	RetrieverContainerCaster() *RetrieverContainer
}

func (s *RetrieverContainer) RetrieverContainerCaster() *RetrieverContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *RetrieverContainer) RRFRetrieverEntryCaster() *RRFRetrieverEntry {
	_ = "STUB: not implemented"
	return nil
}
