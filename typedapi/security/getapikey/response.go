package getapikey

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	ApiKeys []types.ApiKey `json:"api_keys"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
