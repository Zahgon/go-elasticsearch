package getdatafeedstats

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Count     int64                 `json:"count"`
	Datafeeds []types.DatafeedStats `json:"datafeeds"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
