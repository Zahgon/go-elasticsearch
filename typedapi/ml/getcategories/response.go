package getcategories

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Categories []types.Category `json:"categories"`
	Count      int64            `json:"count"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
