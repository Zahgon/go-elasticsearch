package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeseriesmetrictype"
)

type FieldCapability struct {
	Aggregatable bool `json:"aggregatable"`

	Indices []string `json:"indices,omitempty"`

	Inference *bool `json:"inference,omitempty"`

	Meta Metadata `json:"meta,omitempty"`

	MetadataField *bool `json:"metadata_field,omitempty"`

	MetricConflictsIndices []string `json:"metric_conflicts_indices,omitempty"`

	NonAggregatableIndices []string `json:"non_aggregatable_indices,omitempty"`

	NonDimensionIndices []string `json:"non_dimension_indices,omitempty"`

	NonInferenceIndices []string `json:"non_inference_indices,omitempty"`

	NonSearchableIndices []string `json:"non_searchable_indices,omitempty"`

	Searchable bool `json:"searchable"`

	TimeSeriesDimension *bool `json:"time_series_dimension,omitempty"`

	TimeSeriesMetric *timeseriesmetrictype.TimeSeriesMetricType `json:"time_series_metric,omitempty"`
	Type             string                                     `json:"type"`
}

func (s *FieldCapability) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewFieldCapability() *FieldCapability { _ = "STUB: not implemented"; return nil }
