package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _topHitsAggregation struct {
	v *types.TopHitsAggregation
}

func NewTopHitsAggregation() *_topHitsAggregation { _ = "STUB: not implemented"; return nil }

func (s *_topHitsAggregation) DocvalueFields(docvaluefields ...types.FieldAndFormatVariant) *_topHitsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_topHitsAggregation) DocvalueFieldsValues(docvaluefieldsvalues []types.FieldAndFormat) *_topHitsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_topHitsAggregation) Explain(explain bool) *_topHitsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_topHitsAggregation) Fields(fields ...types.FieldAndFormatVariant) *_topHitsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_topHitsAggregation) FieldsValues(fieldsvalues []types.FieldAndFormat) *_topHitsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_topHitsAggregation) From(from int) *_topHitsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_topHitsAggregation) Highlight(highlight types.HighlightVariant) *_topHitsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_topHitsAggregation) ScriptFields(scriptfields map[string]types.ScriptField) *_topHitsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_topHitsAggregation) AddScriptField(key string, value types.ScriptFieldVariant) *_topHitsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_topHitsAggregation) SeqNoPrimaryTerm(seqnoprimaryterm bool) *_topHitsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_topHitsAggregation) Size(size int) *_topHitsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_topHitsAggregation) Sort(sorts ...types.SortCombinationsVariant) *_topHitsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_topHitsAggregation) SortValues(sortvalues []types.SortCombinations) *_topHitsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_topHitsAggregation) Source_(sourceconfig types.SourceConfigVariant) *_topHitsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_topHitsAggregation) StoredFields(fields ...string) *_topHitsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_topHitsAggregation) TrackScores(trackscores bool) *_topHitsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_topHitsAggregation) Version(version bool) *_topHitsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_topHitsAggregation) Field(field string) *_topHitsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_topHitsAggregation) Missing(missing types.MissingVariant) *_topHitsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_topHitsAggregation) Script(script types.ScriptVariant) *_topHitsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_topHitsAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_topHitsAggregation) TopHitsAggregationCaster() *types.TopHitsAggregation {
	_ = "STUB: not implemented"
	return nil
}
