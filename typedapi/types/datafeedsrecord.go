package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/datafeedstate"
)

type DatafeedsRecord struct {
	AssignmentExplanation *string `json:"assignment_explanation,omitempty"`

	BucketsCount *string `json:"buckets.count,omitempty"`

	Id *string `json:"id,omitempty"`

	NodeAddress *string `json:"node.address,omitempty"`

	NodeEphemeralId *string `json:"node.ephemeral_id,omitempty"`

	NodeId *string `json:"node.id,omitempty"`

	NodeName *string `json:"node.name,omitempty"`

	SearchBucketAvg *string `json:"search.bucket_avg,omitempty"`

	SearchCount *string `json:"search.count,omitempty"`

	SearchExpAvgHour *string `json:"search.exp_avg_hour,omitempty"`

	SearchTime *string `json:"search.time,omitempty"`

	State *datafeedstate.DatafeedState `json:"state,omitempty"`
}

func (s *DatafeedsRecord) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewDatafeedsRecord() *DatafeedsRecord { _ = "STUB: not implemented"; return nil }
