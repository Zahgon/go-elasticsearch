package types

import (
	"encoding/json"
)

type RankEvalRequestItem struct {
	Id string `json:"id"`

	Params map[string]json.RawMessage `json:"params,omitempty"`

	Ratings []DocumentRating `json:"ratings"`

	Request *RankEvalQuery `json:"request,omitempty"`

	TemplateId *string `json:"template_id,omitempty"`
}

func (s *RankEvalRequestItem) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewRankEvalRequestItem() *RankEvalRequestItem { _ = "STUB: not implemented"; return nil }

type RankEvalRequestItemVariant interface {
	RankEvalRequestItemCaster() *RankEvalRequestItem
}

func (s *RankEvalRequestItem) RankEvalRequestItemCaster() *RankEvalRequestItem {
	_ = "STUB: not implemented"
	return nil
}
