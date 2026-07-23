package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _aggregations struct {
	v *types.Aggregations
}

func NewAggregations() *_aggregations { _ = "STUB: not implemented"; return nil }

func (s *_aggregations) AdditionalAggregationsProperty(key string, value json.RawMessage) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) AdjacencyMatrix(adjacencymatrix types.AdjacencyMatrixAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) Aggregations(aggregations map[string]types.Aggregations) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) AddAggregation(key string, value types.AggregationsVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) AutoDateHistogram(autodatehistogram types.AutoDateHistogramAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) Avg(avg types.AverageAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) AvgBucket(avgbucket types.AverageBucketAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) Boxplot(boxplot types.BoxplotAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) BucketCorrelation(bucketcorrelation types.BucketCorrelationAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) BucketCountKsTest(bucketcountkstest types.BucketKsAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) BucketScript(bucketscript types.BucketScriptAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) BucketSelector(bucketselector types.BucketSelectorAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) BucketSort(bucketsort types.BucketSortAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) Cardinality(cardinality types.CardinalityAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) CartesianBounds(cartesianbounds types.CartesianBoundsAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) CartesianCentroid(cartesiancentroid types.CartesianCentroidAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) CategorizeText(categorizetext types.CategorizeTextAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) ChangePoint(changepoint types.ChangePointAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) Children(children types.ChildrenAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) Composite(composite types.CompositeAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) CumulativeCardinality(cumulativecardinality types.CumulativeCardinalityAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) CumulativeSum(cumulativesum types.CumulativeSumAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) DateHistogram(datehistogram types.DateHistogramAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) DateRange(daterange types.DateRangeAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) Derivative(derivative types.DerivativeAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) DiversifiedSampler(diversifiedsampler types.DiversifiedSamplerAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) ExtendedStats(extendedstats types.ExtendedStatsAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) ExtendedStatsBucket(extendedstatsbucket types.ExtendedStatsBucketAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) Filter(filter types.QueryVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) Filters(filters types.FiltersAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) FrequentItemSets(frequentitemsets types.FrequentItemSetsAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) GeoBounds(geobounds types.GeoBoundsAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) GeoCentroid(geocentroid types.GeoCentroidAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) GeoDistance(geodistance types.GeoDistanceAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) GeoLine(geoline types.GeoLineAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) GeohashGrid(geohashgrid types.GeoHashGridAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) GeohexGrid(geohexgrid types.GeohexGridAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) GeotileGrid(geotilegrid types.GeoTileGridAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) Global(global types.GlobalAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) Histogram(histogram types.HistogramAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) Inference(inference types.InferenceAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) IpPrefix(ipprefix types.IpPrefixAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) IpRange(iprange types.IpRangeAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) Line(line types.GeoLineAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) MatrixStats(matrixstats types.MatrixStatsAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) Max(max types.MaxAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) MaxBucket(maxbucket types.MaxBucketAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) MedianAbsoluteDeviation(medianabsolutedeviation types.MedianAbsoluteDeviationAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) Meta(metadata types.MetadataVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) Min(min types.MinAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) MinBucket(minbucket types.MinBucketAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) Missing(missing types.MissingAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) MovingAvg(movingaverageaggregation types.MovingAverageAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) MovingFn(movingfn types.MovingFunctionAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) MovingPercentiles(movingpercentiles types.MovingPercentilesAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) MultiTerms(multiterms types.MultiTermsAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) Nested(nested types.NestedAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) Normalize(normalize types.NormalizeAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) Parent(parent types.ParentAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) PercentileRanks(percentileranks types.PercentileRanksAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) Percentiles(percentiles types.PercentilesAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) PercentilesBucket(percentilesbucket types.PercentilesBucketAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) RandomSampler(randomsampler types.RandomSamplerAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) Range(range_ types.RangeAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) RareTerms(rareterms types.RareTermsAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) Rate(rate types.RateAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) ReverseNested(reversenested types.ReverseNestedAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) Sampler(sampler types.SamplerAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) ScriptedMetric(scriptedmetric types.ScriptedMetricAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) SerialDiff(serialdiff types.SerialDifferencingAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) SignificantTerms(significantterms types.SignificantTermsAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) SignificantText(significanttext types.SignificantTextAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) Stats(stats types.StatsAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) StatsBucket(statsbucket types.StatsBucketAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) StringStats(stringstats types.StringStatsAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) Sum(sum types.SumAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) SumBucket(sumbucket types.SumBucketAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) TTest(ttest types.TTestAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) Terms(terms types.TermsAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) TimeSeries(timeseries types.TimeSeriesAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) TopHits(tophits types.TopHitsAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) TopMetrics(topmetrics types.TopMetricsAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) ValueCount(valuecount types.ValueCountAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) VariableWidthHistogram(variablewidthhistogram types.VariableWidthHistogramAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) WeightedAvg(weightedavg types.WeightedAverageAggregationVariant) *_aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregations) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}
