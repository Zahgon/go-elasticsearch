package getfeatureupgradestatus

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/migrationstatus"
)

type Response struct {
	Features        []types.GetMigrationFeature     `json:"features"`
	MigrationStatus migrationstatus.MigrationStatus `json:"migration_status"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
