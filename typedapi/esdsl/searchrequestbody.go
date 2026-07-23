package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _searchRequestBody struct {
	v *types.SearchRequestBody
}

func NewSearchRequestBody() *_searchRequestBody { _ = "STUB: not implemented"; return nil }

func (s *_searchRequestBody) Aggregations(aggregations map[string]types.Aggregations) *_searchRequestBody {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchRequestBody) AddAggregation(key string, value types.AggregationsVariant) *_searchRequestBody {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchRequestBody) Collapse(collapse types.FieldCollapseVariant) *_searchRequestBody {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchRequestBody) DocvalueFields(docvaluefields ...types.FieldAndFormatVariant) *_searchRequestBody {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchRequestBody) DocvalueFieldsValues(docvaluefieldsvalues []types.FieldAndFormat) *_searchRequestBody {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchRequestBody) Explain(explain bool) *_searchRequestBody {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchRequestBody) Ext(ext map[string]json.RawMessage) *_searchRequestBody {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchRequestBody) AddExt(key string, value json.RawMessage) *_searchRequestBody {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchRequestBody) Fields(fields ...types.FieldAndFormatVariant) *_searchRequestBody {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchRequestBody) FieldsValues(fieldsvalues []types.FieldAndFormat) *_searchRequestBody {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchRequestBody) From(from int) *_searchRequestBody {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchRequestBody) Highlight(highlight types.HighlightVariant) *_searchRequestBody {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchRequestBody) IndicesBoost(indicesboost []map[string]types.Float64) *_searchRequestBody {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchRequestBody) Knn(knns ...types.KnnSearchVariant) *_searchRequestBody {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchRequestBody) MinScore(minscore types.Float64) *_searchRequestBody {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchRequestBody) Pit(pit types.PointInTimeReferenceVariant) *_searchRequestBody {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchRequestBody) PostFilter(postfilter types.QueryVariant) *_searchRequestBody {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchRequestBody) Profile(profile bool) *_searchRequestBody {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchRequestBody) Query(query types.QueryVariant) *_searchRequestBody {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchRequestBody) Rank(rank types.RankContainerVariant) *_searchRequestBody {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchRequestBody) Rescore(rescores ...types.RescoreVariant) *_searchRequestBody {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchRequestBody) Retriever(retriever types.RetrieverContainerVariant) *_searchRequestBody {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchRequestBody) RuntimeMappings(runtimefields types.RuntimeFieldsVariant) *_searchRequestBody {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchRequestBody) ScriptFields(scriptfields map[string]types.ScriptField) *_searchRequestBody {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchRequestBody) AddScriptField(key string, value types.ScriptFieldVariant) *_searchRequestBody {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchRequestBody) SearchAfter(sortresults ...types.FieldValueVariant) *_searchRequestBody {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchRequestBody) SearchAfterValues(sortresultsvalues []types.FieldValue) *_searchRequestBody {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchRequestBody) SeqNoPrimaryTerm(seqnoprimaryterm bool) *_searchRequestBody {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchRequestBody) Size(size int) *_searchRequestBody {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchRequestBody) Slice(slice types.SlicedScrollVariant) *_searchRequestBody {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchRequestBody) Sort(sorts ...types.SortCombinationsVariant) *_searchRequestBody {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchRequestBody) SortValues(sortvalues []types.SortCombinations) *_searchRequestBody {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchRequestBody) Source_(sourceconfig types.SourceConfigVariant) *_searchRequestBody {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchRequestBody) Stats(stats ...string) *_searchRequestBody {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchRequestBody) StoredFields(fields ...string) *_searchRequestBody {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchRequestBody) Suggest(suggest types.SuggesterVariant) *_searchRequestBody {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchRequestBody) TerminateAfter(terminateafter int64) *_searchRequestBody {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchRequestBody) Timeout(timeout string) *_searchRequestBody {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchRequestBody) TrackScores(trackscores bool) *_searchRequestBody {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchRequestBody) TrackTotalHits(trackhits types.TrackHitsVariant) *_searchRequestBody {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchRequestBody) Version(version bool) *_searchRequestBody {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchRequestBody) SearchRequestBodyCaster() *types.SearchRequestBody {
	_ = "STUB: not implemented"
	return nil
}
