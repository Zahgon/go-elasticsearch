package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _rankEvalRequestItem struct {
	v *types.RankEvalRequestItem
}

func NewRankEvalRequestItem() *_rankEvalRequestItem { _ = "STUB: not implemented"; return nil }

func (s *_rankEvalRequestItem) Id(id string) *_rankEvalRequestItem {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankEvalRequestItem) Params(params map[string]json.RawMessage) *_rankEvalRequestItem {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankEvalRequestItem) AddParam(key string, value json.RawMessage) *_rankEvalRequestItem {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankEvalRequestItem) Ratings(ratings ...types.DocumentRatingVariant) *_rankEvalRequestItem {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankEvalRequestItem) RatingsValues(ratingsvalues []types.DocumentRating) *_rankEvalRequestItem {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankEvalRequestItem) Request(request types.RankEvalQueryVariant) *_rankEvalRequestItem {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankEvalRequestItem) TemplateId(id string) *_rankEvalRequestItem {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankEvalRequestItem) RankEvalRequestItemCaster() *types.RankEvalRequestItem {
	_ = "STUB: not implemented"
	return nil
}
