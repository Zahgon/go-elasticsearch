package getsettings

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Index types.IndexSettings `json:"index"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
