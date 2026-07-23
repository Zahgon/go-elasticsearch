package putprivileges

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request = map[string]map[string]types.PrivilegesActions

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }
