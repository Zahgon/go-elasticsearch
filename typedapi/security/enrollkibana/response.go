package enrollkibana

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	HttpCa string            `json:"http_ca"`
	Token  types.KibanaToken `json:"token"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
