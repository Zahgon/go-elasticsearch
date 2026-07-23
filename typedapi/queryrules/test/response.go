package test

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	MatchedRules      []types.QueryRulesetMatchedRule `json:"matched_rules"`
	TotalMatchedRules int                             `json:"total_matched_rules"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
