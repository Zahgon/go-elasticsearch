package index

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/failurestorestatus"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/result"
)

type Response struct {
	FailureStore  *failurestorestatus.FailureStoreStatus `json:"failure_store,omitempty"`
	ForcedRefresh *bool                                  `json:"forced_refresh,omitempty"`

	Id_ string `json:"_id"`

	Index_ string `json:"_index"`

	PrimaryTerm_ *int64 `json:"_primary_term,omitempty"`

	Result result.Result `json:"result"`

	SeqNo_ *int64 `json:"_seq_no,omitempty"`

	Shards_ types.ShardStatistics `json:"_shards"`

	Version_ int64 `json:"_version"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
