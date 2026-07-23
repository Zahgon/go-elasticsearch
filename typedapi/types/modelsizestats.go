package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/categorizationstatus"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/memorystatus"
)

type ModelSizeStats struct {
	AssignmentMemoryBasis         *string                                   `json:"assignment_memory_basis,omitempty"`
	BucketAllocationFailuresCount int64                                     `json:"bucket_allocation_failures_count"`
	CategorizationStatus          categorizationstatus.CategorizationStatus `json:"categorization_status"`
	CategorizedDocCount           int                                       `json:"categorized_doc_count"`
	DeadCategoryCount             int                                       `json:"dead_category_count"`
	FailedCategoryCount           int                                       `json:"failed_category_count"`
	FrequentCategoryCount         int                                       `json:"frequent_category_count"`
	JobId                         string                                    `json:"job_id"`
	LogTime                       DateTime                                  `json:"log_time"`
	MemoryStatus                  memorystatus.MemoryStatus                 `json:"memory_status"`
	ModelBytes                    ByteSize                                  `json:"model_bytes"`
	ModelBytesExceeded            ByteSize                                  `json:"model_bytes_exceeded,omitempty"`
	ModelBytesMemoryLimit         ByteSize                                  `json:"model_bytes_memory_limit,omitempty"`
	OutputMemoryAllocatorBytes    ByteSize                                  `json:"output_memory_allocator_bytes,omitempty"`
	PeakModelBytes                ByteSize                                  `json:"peak_model_bytes,omitempty"`
	RareCategoryCount             int                                       `json:"rare_category_count"`
	ResultType                    string                                    `json:"result_type"`
	Timestamp                     *int64                                    `json:"timestamp,omitempty"`
	TotalByFieldCount             int64                                     `json:"total_by_field_count"`
	TotalCategoryCount            int                                       `json:"total_category_count"`
	TotalOverFieldCount           int64                                     `json:"total_over_field_count"`
	TotalPartitionFieldCount      int64                                     `json:"total_partition_field_count"`
}

func (s *ModelSizeStats) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewModelSizeStats() *ModelSizeStats { _ = "STUB: not implemented"; return nil }
