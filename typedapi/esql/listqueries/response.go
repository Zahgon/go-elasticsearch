package listqueries

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Queries map[string]types.Body `json:"queries"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
