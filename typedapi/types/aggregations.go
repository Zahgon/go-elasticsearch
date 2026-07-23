package types

import (
	"encoding/json"
)

type Aggregations struct {
	AdditionalAggregationsProperty map[string]json.RawMessage `json:"-"`

	AdjacencyMatrix *AdjacencyMatrixAggregation `json:"adjacency_matrix,omitempty"`

	Aggregations map[string]Aggregations `json:"aggregations,omitempty"`

	AutoDateHistogram *AutoDateHistogramAggregation `json:"auto_date_histogram,omitempty"`

	Avg *AverageAggregation `json:"avg,omitempty"`

	AvgBucket *AverageBucketAggregation `json:"avg_bucket,omitempty"`

	Boxplot *BoxplotAggregation `json:"boxplot,omitempty"`

	BucketCorrelation *BucketCorrelationAggregation `json:"bucket_correlation,omitempty"`

	BucketCountKsTest *BucketKsAggregation `json:"bucket_count_ks_test,omitempty"`

	BucketScript *BucketScriptAggregation `json:"bucket_script,omitempty"`

	BucketSelector *BucketSelectorAggregation `json:"bucket_selector,omitempty"`

	BucketSort *BucketSortAggregation `json:"bucket_sort,omitempty"`

	Cardinality *CardinalityAggregation `json:"cardinality,omitempty"`

	CartesianBounds *CartesianBoundsAggregation `json:"cartesian_bounds,omitempty"`

	CartesianCentroid *CartesianCentroidAggregation `json:"cartesian_centroid,omitempty"`

	CategorizeText *CategorizeTextAggregation `json:"categorize_text,omitempty"`

	ChangePoint *ChangePointAggregation `json:"change_point,omitempty"`

	Children *ChildrenAggregation `json:"children,omitempty"`

	Composite *CompositeAggregation `json:"composite,omitempty"`

	CumulativeCardinality *CumulativeCardinalityAggregation `json:"cumulative_cardinality,omitempty"`

	CumulativeSum *CumulativeSumAggregation `json:"cumulative_sum,omitempty"`

	DateHistogram *DateHistogramAggregation `json:"date_histogram,omitempty"`

	DateRange *DateRangeAggregation `json:"date_range,omitempty"`

	Derivative *DerivativeAggregation `json:"derivative,omitempty"`

	DiversifiedSampler *DiversifiedSamplerAggregation `json:"diversified_sampler,omitempty"`

	ExtendedStats *ExtendedStatsAggregation `json:"extended_stats,omitempty"`

	ExtendedStatsBucket *ExtendedStatsBucketAggregation `json:"extended_stats_bucket,omitempty"`

	Filter *Query `json:"filter,omitempty"`

	Filters *FiltersAggregation `json:"filters,omitempty"`

	FrequentItemSets *FrequentItemSetsAggregation `json:"frequent_item_sets,omitempty"`

	GeoBounds *GeoBoundsAggregation `json:"geo_bounds,omitempty"`

	GeoCentroid *GeoCentroidAggregation `json:"geo_centroid,omitempty"`

	GeoDistance *GeoDistanceAggregation `json:"geo_distance,omitempty"`

	GeoLine *GeoLineAggregation `json:"geo_line,omitempty"`

	GeohashGrid *GeoHashGridAggregation `json:"geohash_grid,omitempty"`

	GeohexGrid *GeohexGridAggregation `json:"geohex_grid,omitempty"`

	GeotileGrid *GeoTileGridAggregation `json:"geotile_grid,omitempty"`

	Global *GlobalAggregation `json:"global,omitempty"`

	Histogram *HistogramAggregation `json:"histogram,omitempty"`

	Inference *InferenceAggregation `json:"inference,omitempty"`

	IpPrefix *IpPrefixAggregation `json:"ip_prefix,omitempty"`

	IpRange *IpRangeAggregation `json:"ip_range,omitempty"`
	Line    *GeoLineAggregation `json:"line,omitempty"`

	MatrixStats *MatrixStatsAggregation `json:"matrix_stats,omitempty"`

	Max *MaxAggregation `json:"max,omitempty"`

	MaxBucket *MaxBucketAggregation `json:"max_bucket,omitempty"`

	MedianAbsoluteDeviation *MedianAbsoluteDeviationAggregation `json:"median_absolute_deviation,omitempty"`
	Meta                    Metadata                            `json:"meta,omitempty"`

	Min *MinAggregation `json:"min,omitempty"`

	MinBucket *MinBucketAggregation `json:"min_bucket,omitempty"`

	Missing   *MissingAggregation      `json:"missing,omitempty"`
	MovingAvg MovingAverageAggregation `json:"moving_avg,omitempty"`

	MovingFn *MovingFunctionAggregation `json:"moving_fn,omitempty"`

	MovingPercentiles *MovingPercentilesAggregation `json:"moving_percentiles,omitempty"`

	MultiTerms *MultiTermsAggregation `json:"multi_terms,omitempty"`

	Nested *NestedAggregation `json:"nested,omitempty"`

	Normalize *NormalizeAggregation `json:"normalize,omitempty"`

	Parent *ParentAggregation `json:"parent,omitempty"`

	PercentileRanks *PercentileRanksAggregation `json:"percentile_ranks,omitempty"`

	Percentiles *PercentilesAggregation `json:"percentiles,omitempty"`

	PercentilesBucket *PercentilesBucketAggregation `json:"percentiles_bucket,omitempty"`

	RandomSampler *RandomSamplerAggregation `json:"random_sampler,omitempty"`

	Range *RangeAggregation `json:"range,omitempty"`

	RareTerms *RareTermsAggregation `json:"rare_terms,omitempty"`

	Rate *RateAggregation `json:"rate,omitempty"`

	ReverseNested *ReverseNestedAggregation `json:"reverse_nested,omitempty"`

	Sampler *SamplerAggregation `json:"sampler,omitempty"`

	ScriptedMetric *ScriptedMetricAggregation `json:"scripted_metric,omitempty"`

	SerialDiff *SerialDifferencingAggregation `json:"serial_diff,omitempty"`

	SignificantTerms *SignificantTermsAggregation `json:"significant_terms,omitempty"`

	SignificantText *SignificantTextAggregation `json:"significant_text,omitempty"`

	Stats *StatsAggregation `json:"stats,omitempty"`

	StatsBucket *StatsBucketAggregation `json:"stats_bucket,omitempty"`

	StringStats *StringStatsAggregation `json:"string_stats,omitempty"`

	Sum *SumAggregation `json:"sum,omitempty"`

	SumBucket *SumBucketAggregation `json:"sum_bucket,omitempty"`

	TTest *TTestAggregation `json:"t_test,omitempty"`

	Terms *TermsAggregation `json:"terms,omitempty"`

	TimeSeries *TimeSeriesAggregation `json:"time_series,omitempty"`

	TopHits *TopHitsAggregation `json:"top_hits,omitempty"`

	TopMetrics *TopMetricsAggregation `json:"top_metrics,omitempty"`

	ValueCount *ValueCountAggregation `json:"value_count,omitempty"`

	VariableWidthHistogram *VariableWidthHistogramAggregation `json:"variable_width_histogram,omitempty"`

	WeightedAvg *WeightedAverageAggregation `json:"weighted_avg,omitempty"`
}

func (s *Aggregations) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s Aggregations) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewAggregations() *Aggregations { _ = "STUB: not implemented"; return nil }

type AggregationsVariant interface {
	AggregationsCaster() *Aggregations
}

func (s *Aggregations) AggregationsCaster() *Aggregations { _ = "STUB: not implemented"; return nil }
