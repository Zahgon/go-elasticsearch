package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/scorenormalizer"
)

type InnerRetriever struct {
	Normalizer scorenormalizer.ScoreNormalizer `json:"normalizer"`
	Retriever  RetrieverContainer              `json:"retriever"`
	Weight     float32                         `json:"weight"`
}

func (s *InnerRetriever) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewInnerRetriever() *InnerRetriever { _ = "STUB: not implemented"; return nil }

type InnerRetrieverVariant interface {
	InnerRetrieverCaster() *InnerRetriever
}

func (s *InnerRetriever) InnerRetrieverCaster() *InnerRetriever {
	_ = "STUB: not implemented"
	return nil
}
