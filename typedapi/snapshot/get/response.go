package get

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Next *string `json:"next,omitempty"`

	Remaining int                          `json:"remaining"`
	Responses []types.SnapshotResponseItem `json:"responses,omitempty"`
	Snapshots []types.SnapshotInfo         `json:"snapshots,omitempty"`

	Total int `json:"total"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
