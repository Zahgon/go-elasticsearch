package types

type RescorerRetriever struct {
	Filter []Query `json:"filter,omitempty"`

	MinScore *float32 `json:"min_score,omitempty"`

	Name_   *string   `json:"_name,omitempty"`
	Rescore []Rescore `json:"rescore"`

	Retriever RetrieverContainer `json:"retriever"`
}

func (s *RescorerRetriever) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRescorerRetriever() *RescorerRetriever { _ = "STUB: not implemented"; return nil }

type RescorerRetrieverVariant interface {
	RescorerRetrieverCaster() *RescorerRetriever
}

func (s *RescorerRetriever) RescorerRetrieverCaster() *RescorerRetriever {
	_ = "STUB: not implemented"
	return nil
}
