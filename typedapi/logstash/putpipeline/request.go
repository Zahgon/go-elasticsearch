package putpipeline

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request = types.LogstashPipeline

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }
