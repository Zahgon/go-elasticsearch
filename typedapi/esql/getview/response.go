package getview

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Views []types.ESQLView `json:"views"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
