package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/missingorder"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/termsaggregationcollectmode"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/termsaggregationexecutionhint"
)

type _termsAggregation struct {
	v *types.TermsAggregation
}

func NewTermsAggregation() *_termsAggregation { _ = "STUB: not implemented"; return nil }

func (s *_termsAggregation) CollectMode(collectmode termsaggregationcollectmode.TermsAggregationCollectMode) *_termsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termsAggregation) Exclude(termsexcludes ...string) *_termsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termsAggregation) ExecutionHint(executionhint termsaggregationexecutionhint.TermsAggregationExecutionHint) *_termsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termsAggregation) Field(field string) *_termsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termsAggregation) Format(format string) *_termsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termsAggregation) Include(termsinclude types.TermsIncludeVariant) *_termsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termsAggregation) MinDocCount(mindoccount int) *_termsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termsAggregation) Missing(missing types.MissingVariant) *_termsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termsAggregation) MissingBucket(missingbucket bool) *_termsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termsAggregation) MissingOrder(missingorder missingorder.MissingOrder) *_termsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termsAggregation) Order(aggregateorder types.AggregateOrderVariant) *_termsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termsAggregation) Script(script types.ScriptVariant) *_termsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termsAggregation) ShardMinDocCount(shardmindoccount int64) *_termsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termsAggregation) ShardSize(shardsize int) *_termsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termsAggregation) ShowTermDocCountError(showtermdoccounterror bool) *_termsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termsAggregation) Size(size int) *_termsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termsAggregation) ValueType(valuetype string) *_termsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termsAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termsAggregation) ApiKeyAggregationContainerCaster() *types.ApiKeyAggregationContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termsAggregation) PivotGroupByContainerCaster() *types.PivotGroupByContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termsAggregation) TermsAggregationCaster() *types.TermsAggregation {
	_ = "STUB: not implemented"
	return nil
}
