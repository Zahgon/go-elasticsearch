//go:generate easyjson $GOFILE

package model

import (
	"github.com/elastic/go-elasticsearch/v9/esutil"
)

type BulkIndexerResponse struct {
	esutil.BulkIndexerResponse
}
