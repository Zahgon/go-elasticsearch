package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _query struct {
	v *types.Query
}

func NewQuery() *_query { _ = "STUB: not implemented"; return nil }

func (s *_query) AdditionalQueryProperty(key string, value json.RawMessage) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) Bool(bool types.BoolQueryVariant) *_query { _ = "STUB: not implemented"; return nil }

func (s *_query) Boosting(boosting types.BoostingQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) CombinedFields(combinedfields types.CombinedFieldsQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) Common(key string, value types.CommonTermsQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) ConstantScore(constantscore types.ConstantScoreQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) DisMax(dismax types.DisMaxQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) DistanceFeature(distancefeaturequery types.DistanceFeatureQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) Exists(exists types.ExistsQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) FunctionScore(functionscore types.FunctionScoreQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) Fuzzy(key string, value types.FuzzyQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) GeoBoundingBox(geoboundingbox types.GeoBoundingBoxQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) GeoDistance(geodistance types.GeoDistanceQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) GeoGrid(key string, value types.GeoGridQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) GeoPolygon(geopolygon types.GeoPolygonQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) GeoShape(geoshape types.GeoShapeQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) HasChild(haschild types.HasChildQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) HasParent(hasparent types.HasParentQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) Ids(ids types.IdsQueryVariant) *_query { _ = "STUB: not implemented"; return nil }

func (s *_query) Interval(key string, value types.IntervalsQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) Knn(knn types.KnnQueryVariant) *_query { _ = "STUB: not implemented"; return nil }

func (s *_query) Match(key string, value types.MatchQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) MatchAll(matchall types.MatchAllQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) MatchBoolPrefix(key string, value types.MatchBoolPrefixQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) MatchNone(matchnone types.MatchNoneQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) MatchPhrase(key string, value types.MatchPhraseQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) MatchPhrasePrefix(key string, value types.MatchPhrasePrefixQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) MoreLikeThis(morelikethis types.MoreLikeThisQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) MultiMatch(multimatch types.MultiMatchQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) Nested(nested types.NestedQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) ParentId(parentid types.ParentIdQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) Percolate(percolate types.PercolateQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) Pinned(pinned types.PinnedQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) Prefix(key string, value types.PrefixQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) QueryString(querystring types.QueryStringQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) Range(key string, value types.RangeQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) RankFeature(rankfeature types.RankFeatureQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) Regexp(key string, value types.RegexpQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) Rule(rule types.RuleQueryVariant) *_query { _ = "STUB: not implemented"; return nil }

func (s *_query) Script(script types.ScriptQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) ScriptScore(scriptscore types.ScriptScoreQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) Semantic(semantic types.SemanticQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) Shape(shape types.ShapeQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) SimpleQueryString(simplequerystring types.SimpleQueryStringQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) SpanContaining(spancontaining types.SpanContainingQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) SpanFieldMasking(spanfieldmasking types.SpanFieldMaskingQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) SpanFirst(spanfirst types.SpanFirstQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) SpanMulti(spanmulti types.SpanMultiTermQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) SpanNear(spannear types.SpanNearQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) SpanNot(spannot types.SpanNotQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) SpanOr(spanor types.SpanOrQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) SpanTerm(key string, value types.SpanTermQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) SpanWithin(spanwithin types.SpanWithinQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) SparseVector(sparsevector types.SparseVectorQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) Term(key string, value types.TermQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) Terms(terms types.TermsQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) TermsSet(key string, value types.TermsSetQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) TextExpansion(key string, value types.TextExpansionQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) Type(type_ types.TypeQueryVariant) *_query { _ = "STUB: not implemented"; return nil }

func (s *_query) WeightedToken(key string, value types.WeightedTokensQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) Wildcard(key string, value types.WildcardQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) Wrapper(wrapper types.WrapperQueryVariant) *_query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_query) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }
