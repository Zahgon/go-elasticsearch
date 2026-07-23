package executewatch

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Id_ string `json:"_id"`

	WatchRecord types.WatchRecord `json:"watch_record"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
