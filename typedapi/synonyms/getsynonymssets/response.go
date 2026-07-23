package getsynonymssets

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Count int `json:"count"`

	Results []types.SynonymsSetItem `json:"results"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
