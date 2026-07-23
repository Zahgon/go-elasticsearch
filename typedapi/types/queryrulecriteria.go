package types

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/queryrulecriteriatype"
)

type QueryRuleCriteria struct {
	Metadata *string `json:"metadata,omitempty"`

	Type queryrulecriteriatype.QueryRuleCriteriaType `json:"type"`

	Values []json.RawMessage `json:"values,omitempty"`
}

func (s *QueryRuleCriteria) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewQueryRuleCriteria() *QueryRuleCriteria { _ = "STUB: not implemented"; return nil }

type QueryRuleCriteriaVariant interface {
	QueryRuleCriteriaCaster() *QueryRuleCriteria
}

func (s *QueryRuleCriteria) QueryRuleCriteriaCaster() *QueryRuleCriteria {
	_ = "STUB: not implemented"
	return nil
}
