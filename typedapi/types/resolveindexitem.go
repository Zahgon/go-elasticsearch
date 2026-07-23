package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/indexmode"
)

type ResolveIndexItem struct {
	Aliases    []string             `json:"aliases,omitempty"`
	Attributes []string             `json:"attributes"`
	DataStream *string              `json:"data_stream,omitempty"`
	Mode       *indexmode.IndexMode `json:"mode,omitempty"`
	Name       string               `json:"name"`
}

func (s *ResolveIndexItem) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewResolveIndexItem() *ResolveIndexItem { _ = "STUB: not implemented"; return nil }
