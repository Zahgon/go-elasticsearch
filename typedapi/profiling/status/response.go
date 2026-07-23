package status

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/profilingoperationmode"
)

type Response struct {
	OperationMode profilingoperationmode.ProfilingOperationMode `json:"operation_mode"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
