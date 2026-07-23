package types

type TextSimilarityReranker struct {
	ChunkRescorer *ChunkRescorer `json:"chunk_rescorer,omitempty"`

	Field string `json:"field"`

	Filter []Query `json:"filter,omitempty"`

	InferenceId *string `json:"inference_id,omitempty"`

	InferenceText string `json:"inference_text"`

	MinScore *float32 `json:"min_score,omitempty"`

	Name_ *string `json:"_name,omitempty"`

	RankWindowSize *int `json:"rank_window_size,omitempty"`

	Retriever RetrieverContainer `json:"retriever"`
}

func (s *TextSimilarityReranker) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewTextSimilarityReranker() *TextSimilarityReranker { _ = "STUB: not implemented"; return nil }

type TextSimilarityRerankerVariant interface {
	TextSimilarityRerankerCaster() *TextSimilarityReranker
}

func (s *TextSimilarityReranker) TextSimilarityRerankerCaster() *TextSimilarityReranker {
	_ = "STUB: not implemented"
	return nil
}
