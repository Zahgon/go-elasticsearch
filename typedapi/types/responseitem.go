package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/failurestorestatus"
)

type ResponseItem struct {
	Error         *ErrorCause                            `json:"error,omitempty"`
	FailureStore  *failurestorestatus.FailureStoreStatus `json:"failure_store,omitempty"`
	ForcedRefresh *bool                                  `json:"forced_refresh,omitempty"`
	Get           *InlineGetDictUserDefined              `json:"get,omitempty"`

	Id_ *string `json:"_id,omitempty"`

	Index_ string `json:"_index"`

	PrimaryTerm_ *int64 `json:"_primary_term,omitempty"`

	Result *string `json:"result,omitempty"`

	SeqNo_ *int64 `json:"_seq_no,omitempty"`

	Shards_ *ShardStatistics `json:"_shards,omitempty"`

	Status int `json:"status"`

	Version_ *int64 `json:"_version,omitempty"`
}

func (s *ResponseItem) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewResponseItem() *ResponseItem { _ = "STUB: not implemented"; return nil }
