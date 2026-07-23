package getdatastreamoptions

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	DataStreams []types.DataStreamWithOptions `json:"data_streams"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
