package getiplocationdatabase

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Databases []types.IpLocationDatabaseConfigurationMetadata `json:"databases"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
