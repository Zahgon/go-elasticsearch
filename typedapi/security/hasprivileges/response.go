package hasprivileges

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Application     types.ApplicationsPrivileges `json:"application"`
	Cluster         map[string]bool              `json:"cluster"`
	HasAllRequested bool                         `json:"has_all_requested"`
	Index           map[string]types.Privileges  `json:"index"`
	Username        string                       `json:"username"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
