package bulk

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/operationtype"
)

type Response struct {
	Errors     bool   `json:"errors"`
	IngestTook *int64 `json:"ingest_took,omitempty"`

	Items []map[operationtype.OperationType]types.ResponseItem `json:"items"`

	Took int64 `json:"took"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
