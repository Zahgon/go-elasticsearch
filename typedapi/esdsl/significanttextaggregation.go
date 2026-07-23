package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/termsaggregationexecutionhint"
)

type _significantTextAggregation struct {
	v *types.SignificantTextAggregation
}

func NewSignificantTextAggregation() *_significantTextAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_significantTextAggregation) BackgroundFilter(backgroundfilter types.QueryVariant) *_significantTextAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_significantTextAggregation) ChiSquare(chisquare types.ChiSquareHeuristicVariant) *_significantTextAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_significantTextAggregation) Exclude(termsexcludes ...string) *_significantTextAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_significantTextAggregation) ExecutionHint(executionhint termsaggregationexecutionhint.TermsAggregationExecutionHint) *_significantTextAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_significantTextAggregation) Field(field string) *_significantTextAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_significantTextAggregation) FilterDuplicateText(filterduplicatetext bool) *_significantTextAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_significantTextAggregation) Gnd(gnd types.GoogleNormalizedDistanceHeuristicVariant) *_significantTextAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_significantTextAggregation) Include(termsinclude types.TermsIncludeVariant) *_significantTextAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_significantTextAggregation) Jlh(jlh types.EmptyObjectVariant) *_significantTextAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_significantTextAggregation) MinDocCount(mindoccount int64) *_significantTextAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_significantTextAggregation) MutualInformation(mutualinformation types.MutualInformationHeuristicVariant) *_significantTextAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_significantTextAggregation) Percentage(percentage types.PercentageScoreHeuristicVariant) *_significantTextAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_significantTextAggregation) ScriptHeuristic(scriptheuristic types.ScriptedHeuristicVariant) *_significantTextAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_significantTextAggregation) ShardMinDocCount(shardmindoccount int64) *_significantTextAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_significantTextAggregation) ShardSize(shardsize int) *_significantTextAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_significantTextAggregation) Size(size int) *_significantTextAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_significantTextAggregation) SourceFields(fields ...string) *_significantTextAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_significantTextAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_significantTextAggregation) SignificantTextAggregationCaster() *types.SignificantTextAggregation {
	_ = "STUB: not implemented"
	return nil
}
