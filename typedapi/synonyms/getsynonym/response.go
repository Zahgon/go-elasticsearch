package getsynonym

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Count int `json:"count"`

	SynonymsSet []types.SynonymRuleRead `json:"synonyms_set"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
