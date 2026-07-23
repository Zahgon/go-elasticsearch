package mount

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Snapshot types.MountedSnapshot `json:"snapshot"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
