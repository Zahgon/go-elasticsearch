package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _metadata struct {
	v types.Metadata
}

func NewMetadata(metadata map[string]json.RawMessage) *_metadata {
	_ = "STUB: not implemented"
	return nil
}

func (u *_metadata) MetadataCaster() *types.Metadata { _ = "STUB: not implemented"; return nil }
