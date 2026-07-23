package getrole

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response map[string]types.Role

func NewResponse() Response { _ = "STUB: not implemented"; return *new(Response) }
