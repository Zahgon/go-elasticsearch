package getmigratereindexstatus

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Complete                     bool                     `json:"complete"`
	Errors                       []types.StatusError      `json:"errors"`
	Exception                    *string                  `json:"exception,omitempty"`
	InProgress                   []types.StatusInProgress `json:"in_progress"`
	Pending                      int                      `json:"pending"`
	StartTime                    types.DateTime           `json:"start_time,omitempty"`
	StartTimeMillis              int64                    `json:"start_time_millis"`
	Successes                    int                      `json:"successes"`
	TotalIndicesInDataStream     int                      `json:"total_indices_in_data_stream"`
	TotalIndicesRequiringUpgrade int                      `json:"total_indices_requiring_upgrade"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }

func (s *Response) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
