package getrepository

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response map[string]types.Repository

func NewResponse() Response { _ = "STUB: not implemented"; return *new(Response) }

func (r Response) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
