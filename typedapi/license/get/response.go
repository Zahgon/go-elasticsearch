package get

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	License types.LicenseInformation `json:"license"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
