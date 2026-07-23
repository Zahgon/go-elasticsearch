package createservicetoken

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Created bool               `json:"created"`
	Token   types.ServiceToken `json:"token"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
