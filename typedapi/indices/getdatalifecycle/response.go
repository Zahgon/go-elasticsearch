package getdatalifecycle

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	DataStreams     []types.DataStreamWithLifecycle `json:"data_streams"`
	GlobalRetention types.GlobalRetention           `json:"global_retention"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
