package nodes

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response []types.NodesRecord

func NewResponse() Response { _ = "STUB: not implemented"; return *new(Response) }
