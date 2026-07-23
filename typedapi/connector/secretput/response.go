package secretput

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/result"
)

type Response struct {
	Result result.Result `json:"result"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
