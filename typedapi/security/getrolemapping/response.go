package getrolemapping

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response map[string]types.SecurityRoleMapping

func NewResponse() Response { _ = "STUB: not implemented"; return *new(Response) }
