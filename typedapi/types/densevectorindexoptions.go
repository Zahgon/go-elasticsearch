package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/densevectorindexoptionstype"
)

type DenseVectorIndexOptions struct {
	ConfidenceInterval *float32 `json:"confidence_interval,omitempty"`

	EfConstruction *int `json:"ef_construction,omitempty"`

	FlatIndexThreshold *int `json:"flat_index_threshold,omitempty"`

	M *int `json:"m,omitempty"`

	OnDiskRescore *bool `json:"on_disk_rescore,omitempty"`

	RescoreVector *DenseVectorIndexOptionsRescoreVector `json:"rescore_vector,omitempty"`

	Type densevectorindexoptionstype.DenseVectorIndexOptionsType `json:"type"`
}

func (s *DenseVectorIndexOptions) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDenseVectorIndexOptions() *DenseVectorIndexOptions { _ = "STUB: not implemented"; return nil }

type DenseVectorIndexOptionsVariant interface {
	DenseVectorIndexOptionsCaster() *DenseVectorIndexOptions
}

func (s *DenseVectorIndexOptions) DenseVectorIndexOptionsCaster() *DenseVectorIndexOptions {
	_ = "STUB: not implemented"
	return nil
}
