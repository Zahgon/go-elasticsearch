package putprivileges

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response map[string]map[string]types.CreatedStatus

func NewResponse() Response { _ = "STUB: not implemented"; return *new(Response) }
