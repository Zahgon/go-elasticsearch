package getrule

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/queryruletype"
)

type Response struct {
	Actions types.QueryRuleActions `json:"actions"`

	Criteria []types.QueryRuleCriteria `json:"criteria"`
	Priority *int                      `json:"priority,omitempty"`

	RuleId string `json:"rule_id"`

	Type queryruletype.QueryRuleType `json:"type"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
