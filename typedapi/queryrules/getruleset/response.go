package getruleset

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Rules []types.QueryRule `json:"rules"`

	RulesetId string `json:"ruleset_id"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
