package msearch

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Responses []types.MsearchResponseItem `json:"responses"`
	Took      int64                       `json:"took"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }

func (s *Response) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
