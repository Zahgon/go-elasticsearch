package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/failurestorestatus"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/result"
)

type WriteResponseBase struct {
	FailureStore  *failurestorestatus.FailureStoreStatus `json:"failure_store,omitempty"`
	ForcedRefresh *bool                                  `json:"forced_refresh,omitempty"`

	Id_ string `json:"_id"`

	Index_ string `json:"_index"`

	PrimaryTerm_ *int64 `json:"_primary_term,omitempty"`

	Result result.Result `json:"result"`

	SeqNo_ *int64 `json:"_seq_no,omitempty"`

	Shards_ ShardStatistics `json:"_shards"`

	Version_ int64 `json:"_version"`
}

func (s *WriteResponseBase) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewWriteResponseBase() *WriteResponseBase { _ = "STUB: not implemented"; return nil }
