package getuser

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response map[string]types.User

func NewResponse() Response { _ = "STUB: not implemented"; return *new(Response) }
