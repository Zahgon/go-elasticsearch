package types

type ExtendedRetrieversSearchUsage struct {
	TextSimilarityReranker *ExtendedTextSimilarityRetrieverUsage `json:"text_similarity_reranker,omitempty"`
}

func NewExtendedRetrieversSearchUsage() *ExtendedRetrieversSearchUsage {
	_ = "STUB: not implemented"
	return nil
}
