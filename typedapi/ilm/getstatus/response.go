package getstatus

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/lifecycleoperationmode"
)

type Response struct {
	OperationMode lifecycleoperationmode.LifecycleOperationMode `json:"operation_mode"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
