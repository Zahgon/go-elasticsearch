package put

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/result"
)

type Response struct {
	Id     string        `json:"id"`
	Result result.Result `json:"result"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
