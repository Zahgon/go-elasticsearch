package putdatastreammappings

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	DataStreams []types.UpdatedDataStreamMappings `json:"data_streams"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
