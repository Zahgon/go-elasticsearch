package getrollupindexcaps

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response map[string]types.IndexCapabilities

func NewResponse() Response { _ = "STUB: not implemented"; return *new(Response) }
