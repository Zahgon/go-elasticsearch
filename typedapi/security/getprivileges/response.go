package getprivileges

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response map[string]map[string]types.PrivilegesActions

func NewResponse() Response { _ = "STUB: not implemented"; return *new(Response) }
