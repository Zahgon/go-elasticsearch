package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/termsaggregationexecutionhint"
)

type _significantTermsAggregation struct {
	v *types.SignificantTermsAggregation
}

func NewSignificantTermsAggregation() *_significantTermsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_significantTermsAggregation) BackgroundFilter(backgroundfilter types.QueryVariant) *_significantTermsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_significantTermsAggregation) ChiSquare(chisquare types.ChiSquareHeuristicVariant) *_significantTermsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_significantTermsAggregation) Exclude(termsexcludes ...string) *_significantTermsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_significantTermsAggregation) ExecutionHint(executionhint termsaggregationexecutionhint.TermsAggregationExecutionHint) *_significantTermsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_significantTermsAggregation) Field(field string) *_significantTermsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_significantTermsAggregation) Gnd(gnd types.GoogleNormalizedDistanceHeuristicVariant) *_significantTermsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_significantTermsAggregation) Include(termsinclude types.TermsIncludeVariant) *_significantTermsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_significantTermsAggregation) Jlh(jlh types.EmptyObjectVariant) *_significantTermsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_significantTermsAggregation) MinDocCount(mindoccount int64) *_significantTermsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_significantTermsAggregation) MutualInformation(mutualinformation types.MutualInformationHeuristicVariant) *_significantTermsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_significantTermsAggregation) PValue(pvalue types.PValueHeuristicVariant) *_significantTermsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_significantTermsAggregation) Percentage(percentage types.PercentageScoreHeuristicVariant) *_significantTermsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_significantTermsAggregation) ScriptHeuristic(scriptheuristic types.ScriptedHeuristicVariant) *_significantTermsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_significantTermsAggregation) ShardMinDocCount(shardmindoccount int64) *_significantTermsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_significantTermsAggregation) ShardSize(shardsize int) *_significantTermsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_significantTermsAggregation) Size(size int) *_significantTermsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_significantTermsAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_significantTermsAggregation) SignificantTermsAggregationCaster() *types.SignificantTermsAggregation {
	_ = "STUB: not implemented"
	return nil
}
