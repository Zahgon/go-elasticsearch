package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/categorizationstatus"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/jobstate"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/memorystatus"
)

type JobsRecord struct {
	AssignmentExplanation *string `json:"assignment_explanation,omitempty"`

	BucketsCount *string `json:"buckets.count,omitempty"`

	BucketsTimeExpAvg *string `json:"buckets.time.exp_avg,omitempty"`

	BucketsTimeExpAvgHour *string `json:"buckets.time.exp_avg_hour,omitempty"`

	BucketsTimeMax *string `json:"buckets.time.max,omitempty"`

	BucketsTimeMin *string `json:"buckets.time.min,omitempty"`

	BucketsTimeTotal *string `json:"buckets.time.total,omitempty"`

	DataBuckets *string `json:"data.buckets,omitempty"`

	DataEarliestRecord *string `json:"data.earliest_record,omitempty"`

	DataEmptyBuckets *string `json:"data.empty_buckets,omitempty"`

	DataInputBytes ByteSize `json:"data.input_bytes,omitempty"`

	DataInputFields *string `json:"data.input_fields,omitempty"`

	DataInputRecords *string `json:"data.input_records,omitempty"`

	DataInvalidDates *string `json:"data.invalid_dates,omitempty"`

	DataLast *string `json:"data.last,omitempty"`

	DataLastEmptyBucket *string `json:"data.last_empty_bucket,omitempty"`

	DataLastSparseBucket *string `json:"data.last_sparse_bucket,omitempty"`

	DataLatestRecord *string `json:"data.latest_record,omitempty"`

	DataMissingFields *string `json:"data.missing_fields,omitempty"`

	DataOutOfOrderTimestamps *string `json:"data.out_of_order_timestamps,omitempty"`

	DataProcessedFields *string `json:"data.processed_fields,omitempty"`

	DataProcessedRecords *string `json:"data.processed_records,omitempty"`

	DataSparseBuckets *string `json:"data.sparse_buckets,omitempty"`

	ForecastsMemoryAvg *string `json:"forecasts.memory.avg,omitempty"`

	ForecastsMemoryMax *string `json:"forecasts.memory.max,omitempty"`

	ForecastsMemoryMin *string `json:"forecasts.memory.min,omitempty"`

	ForecastsMemoryTotal *string `json:"forecasts.memory.total,omitempty"`

	ForecastsRecordsAvg *string `json:"forecasts.records.avg,omitempty"`

	ForecastsRecordsMax *string `json:"forecasts.records.max,omitempty"`

	ForecastsRecordsMin *string `json:"forecasts.records.min,omitempty"`

	ForecastsRecordsTotal *string `json:"forecasts.records.total,omitempty"`

	ForecastsTimeAvg *string `json:"forecasts.time.avg,omitempty"`

	ForecastsTimeMax *string `json:"forecasts.time.max,omitempty"`

	ForecastsTimeMin *string `json:"forecasts.time.min,omitempty"`

	ForecastsTimeTotal *string `json:"forecasts.time.total,omitempty"`

	ForecastsTotal *string `json:"forecasts.total,omitempty"`

	Id *string `json:"id,omitempty"`

	ModelBucketAllocationFailures *string `json:"model.bucket_allocation_failures,omitempty"`

	ModelByFields *string `json:"model.by_fields,omitempty"`

	ModelBytes ByteSize `json:"model.bytes,omitempty"`

	ModelBytesExceeded ByteSize `json:"model.bytes_exceeded,omitempty"`

	ModelCategorizationStatus *categorizationstatus.CategorizationStatus `json:"model.categorization_status,omitempty"`

	ModelCategorizedDocCount *string `json:"model.categorized_doc_count,omitempty"`

	ModelDeadCategoryCount *string `json:"model.dead_category_count,omitempty"`

	ModelFailedCategoryCount *string `json:"model.failed_category_count,omitempty"`

	ModelFrequentCategoryCount *string `json:"model.frequent_category_count,omitempty"`

	ModelLogTime *string `json:"model.log_time,omitempty"`

	ModelMemoryLimit *string `json:"model.memory_limit,omitempty"`

	ModelMemoryStatus *memorystatus.MemoryStatus `json:"model.memory_status,omitempty"`

	ModelOverFields *string `json:"model.over_fields,omitempty"`

	ModelPartitionFields *string `json:"model.partition_fields,omitempty"`

	ModelRareCategoryCount *string `json:"model.rare_category_count,omitempty"`

	ModelTimestamp *string `json:"model.timestamp,omitempty"`

	ModelTotalCategoryCount *string `json:"model.total_category_count,omitempty"`

	NodeAddress *string `json:"node.address,omitempty"`

	NodeEphemeralId *string `json:"node.ephemeral_id,omitempty"`

	NodeId *string `json:"node.id,omitempty"`

	NodeName *string `json:"node.name,omitempty"`

	OpenedTime *string `json:"opened_time,omitempty"`

	State *jobstate.JobState `json:"state,omitempty"`
}

func (s *JobsRecord) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewJobsRecord() *JobsRecord { _ = "STUB: not implemented"; return nil }
