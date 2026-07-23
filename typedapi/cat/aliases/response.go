package aliases

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response []types.AliasesRecord

func NewResponse() Response { _ = "STUB: not implemented"; return *new(Response) }
