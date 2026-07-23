package repositoryanalyze

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	BlobCount int `json:"blob_count"`

	BlobPath string `json:"blob_path"`

	Concurrency int `json:"concurrency"`

	CoordinatingNode types.SnapshotNodeInfo `json:"coordinating_node"`

	DeleteElapsed types.Duration `json:"delete_elapsed"`

	DeleteElapsedNanos int64 `json:"delete_elapsed_nanos"`

	Details types.DetailsInfo `json:"details"`

	EarlyReadNodeCount int `json:"early_read_node_count"`

	IssuesDetected []string `json:"issues_detected"`

	ListingElapsed types.Duration `json:"listing_elapsed"`

	ListingElapsedNanos int64 `json:"listing_elapsed_nanos"`

	MaxBlobSize types.ByteSize `json:"max_blob_size"`

	MaxBlobSizeBytes int64 `json:"max_blob_size_bytes"`

	MaxTotalDataSize types.ByteSize `json:"max_total_data_size"`

	MaxTotalDataSizeBytes int64 `json:"max_total_data_size_bytes"`

	RareActionProbability types.Float64 `json:"rare_action_probability"`

	ReadNodeCount int `json:"read_node_count"`

	Repository string `json:"repository"`

	Seed int64 `json:"seed"`

	Summary types.SummaryInfo `json:"summary"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
