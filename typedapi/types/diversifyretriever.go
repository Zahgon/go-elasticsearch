package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/diversifyretrievertypes"
)

type DiversifyRetriever struct {
	Field string `json:"field"`

	Filter []Query `json:"filter,omitempty"`

	Lambda *float32 `json:"lambda,omitempty"`

	MinScore *float32 `json:"min_score,omitempty"`

	Name_ *string `json:"_name,omitempty"`

	QueryVector []float32 `json:"query_vector,omitempty"`

	QueryVectorBuilder *QueryVectorBuilder `json:"query_vector_builder,omitempty"`

	RankWindowSize *int `json:"rank_window_size,omitempty"`

	Retriever RetrieverContainer `json:"retriever"`

	Size *int `json:"size,omitempty"`

	Type diversifyretrievertypes.DiversifyRetrieverTypes `json:"type"`
}

func (s *DiversifyRetriever) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDiversifyRetriever() *DiversifyRetriever { _ = "STUB: not implemented"; return nil }

type DiversifyRetrieverVariant interface {
	DiversifyRetrieverCaster() *DiversifyRetriever
}

func (s *DiversifyRetriever) DiversifyRetrieverCaster() *DiversifyRetriever {
	_ = "STUB: not implemented"
	return nil
}
