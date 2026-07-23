package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/scorenormalizer"
)

type LinearRetriever struct {
	Fields []string `json:"fields,omitempty"`

	Filter []Query `json:"filter,omitempty"`

	MinScore *float32 `json:"min_score,omitempty"`

	Name_          *string                          `json:"_name,omitempty"`
	Normalizer     *scorenormalizer.ScoreNormalizer `json:"normalizer,omitempty"`
	Query          *string                          `json:"query,omitempty"`
	RankWindowSize *int                             `json:"rank_window_size,omitempty"`

	Retrievers []InnerRetriever `json:"retrievers,omitempty"`
}

func (s *LinearRetriever) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewLinearRetriever() *LinearRetriever { _ = "STUB: not implemented"; return nil }

type LinearRetrieverVariant interface {
	LinearRetrieverCaster() *LinearRetriever
}

func (s *LinearRetriever) LinearRetrieverCaster() *LinearRetriever {
	_ = "STUB: not implemented"
	return nil
}
