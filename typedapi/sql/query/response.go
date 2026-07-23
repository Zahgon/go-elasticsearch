package query

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Columns []types.Column `json:"columns,omitempty"`

	Cursor *string `json:"cursor,omitempty"`

	Id *string `json:"id,omitempty"`

	IsPartial *bool `json:"is_partial,omitempty"`

	IsRunning *bool `json:"is_running,omitempty"`

	Rows [][]json.RawMessage `json:"rows"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
