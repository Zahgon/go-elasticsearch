package getmanyrouting

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response map[string]types.ProjectRoutingExpression

func NewResponse() Response { _ = "STUB: not implemented"; return *new(Response) }
