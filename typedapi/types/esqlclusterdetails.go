package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/esqlclusterstatus"
)

type EsqlClusterDetails struct {
	Failures []EsqlShardFailure                  `json:"failures,omitempty"`
	Indices  string                              `json:"indices"`
	Shards_  *EsqlShardInfo                      `json:"_shards,omitempty"`
	Status   esqlclusterstatus.EsqlClusterStatus `json:"status"`
	Took     *int64                              `json:"took,omitempty"`
}

func (s *EsqlClusterDetails) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewEsqlClusterDetails() *EsqlClusterDetails { _ = "STUB: not implemented"; return nil }
