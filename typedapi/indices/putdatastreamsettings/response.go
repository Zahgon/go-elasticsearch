package putdatastreamsettings

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	DataStreams []types.UpdatedDataStreamSettings `json:"data_streams"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
