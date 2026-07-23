package migratereindex

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request = types.MigrateReindex

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }
