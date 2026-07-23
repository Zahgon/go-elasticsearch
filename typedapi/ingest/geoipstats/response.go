package geoipstats

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Nodes map[string]types.GeoIpNodeDatabases `json:"nodes"`

	Stats types.GeoIpDownloadStatistics `json:"stats"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
