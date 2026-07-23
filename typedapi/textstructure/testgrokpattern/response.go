package testgrokpattern

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Matches []types.MatchedText `json:"matches"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
