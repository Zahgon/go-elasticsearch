package info

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Build    types.BuildInformation          `json:"build"`
	Features types.XpackFeatures             `json:"features"`
	License  types.MinimalLicenseInformation `json:"license"`
	Tagline  string                          `json:"tagline"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
