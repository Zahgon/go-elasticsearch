package resolveindex

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Aliases     []types.ResolveIndexAliasItem       `json:"aliases"`
	DataStreams []types.ResolveIndexDataStreamsItem `json:"data_streams"`
	Indices     []types.ResolveIndexItem            `json:"indices"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
