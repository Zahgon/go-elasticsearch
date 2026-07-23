package types

import (
	"encoding/json"
)

type Query struct {
	AdditionalQueryProperty map[string]json.RawMessage `json:"-"`

	Bool *BoolQuery `json:"bool,omitempty"`

	Boosting *BoostingQuery `json:"boosting,omitempty"`

	CombinedFields *CombinedFieldsQuery        `json:"combined_fields,omitempty"`
	Common         map[string]CommonTermsQuery `json:"common,omitempty"`

	ConstantScore *ConstantScoreQuery `json:"constant_score,omitempty"`

	DisMax *DisMaxQuery `json:"dis_max,omitempty"`

	DistanceFeature DistanceFeatureQuery `json:"distance_feature,omitempty"`

	Exists *ExistsQuery `json:"exists,omitempty"`

	FunctionScore *FunctionScoreQuery `json:"function_score,omitempty"`

	Fuzzy map[string]FuzzyQuery `json:"fuzzy,omitempty"`

	GeoBoundingBox *GeoBoundingBoxQuery `json:"geo_bounding_box,omitempty"`

	GeoDistance *GeoDistanceQuery `json:"geo_distance,omitempty"`

	GeoGrid    map[string]GeoGridQuery `json:"geo_grid,omitempty"`
	GeoPolygon *GeoPolygonQuery        `json:"geo_polygon,omitempty"`

	GeoShape *GeoShapeQuery `json:"geo_shape,omitempty"`

	HasChild *HasChildQuery `json:"has_child,omitempty"`

	HasParent *HasParentQuery `json:"has_parent,omitempty"`

	Ids *IdsQuery `json:"ids,omitempty"`

	Intervals map[string]IntervalsQuery `json:"intervals,omitempty"`

	Knn *KnnQuery `json:"knn,omitempty"`

	Match map[string]MatchQuery `json:"match,omitempty"`

	MatchAll *MatchAllQuery `json:"match_all,omitempty"`

	MatchBoolPrefix map[string]MatchBoolPrefixQuery `json:"match_bool_prefix,omitempty"`

	MatchNone *MatchNoneQuery `json:"match_none,omitempty"`

	MatchPhrase map[string]MatchPhraseQuery `json:"match_phrase,omitempty"`

	MatchPhrasePrefix map[string]MatchPhrasePrefixQuery `json:"match_phrase_prefix,omitempty"`

	MoreLikeThis *MoreLikeThisQuery `json:"more_like_this,omitempty"`

	MultiMatch *MultiMatchQuery `json:"multi_match,omitempty"`

	Nested *NestedQuery `json:"nested,omitempty"`

	ParentId *ParentIdQuery `json:"parent_id,omitempty"`

	Percolate *PercolateQuery `json:"percolate,omitempty"`

	Pinned *PinnedQuery `json:"pinned,omitempty"`

	Prefix map[string]PrefixQuery `json:"prefix,omitempty"`

	QueryString *QueryStringQuery `json:"query_string,omitempty"`

	Range map[string]RangeQuery `json:"range,omitempty"`

	RankFeature *RankFeatureQuery `json:"rank_feature,omitempty"`

	Regexp map[string]RegexpQuery `json:"regexp,omitempty"`
	Rule   *RuleQuery             `json:"rule,omitempty"`

	Script *ScriptQuery `json:"script,omitempty"`

	ScriptScore *ScriptScoreQuery `json:"script_score,omitempty"`

	Semantic *SemanticQuery `json:"semantic,omitempty"`

	Shape *ShapeQuery `json:"shape,omitempty"`

	SimpleQueryString *SimpleQueryStringQuery `json:"simple_query_string,omitempty"`

	SpanContaining *SpanContainingQuery `json:"span_containing,omitempty"`

	SpanFieldMasking *SpanFieldMaskingQuery `json:"span_field_masking,omitempty"`

	SpanFirst *SpanFirstQuery `json:"span_first,omitempty"`

	SpanMulti *SpanMultiTermQuery `json:"span_multi,omitempty"`

	SpanNear *SpanNearQuery `json:"span_near,omitempty"`

	SpanNot *SpanNotQuery `json:"span_not,omitempty"`

	SpanOr *SpanOrQuery `json:"span_or,omitempty"`

	SpanTerm map[string]SpanTermQuery `json:"span_term,omitempty"`

	SpanWithin *SpanWithinQuery `json:"span_within,omitempty"`

	SparseVector *SparseVectorQuery `json:"sparse_vector,omitempty"`

	Term map[string]TermQuery `json:"term,omitempty"`

	Terms *TermsQuery `json:"terms,omitempty"`

	TermsSet map[string]TermsSetQuery `json:"terms_set,omitempty"`

	TextExpansion map[string]TextExpansionQuery `json:"text_expansion,omitempty"`
	Type          *TypeQuery                    `json:"type,omitempty"`

	WeightedTokens map[string]WeightedTokensQuery `json:"weighted_tokens,omitempty"`

	Wildcard map[string]WildcardQuery `json:"wildcard,omitempty"`

	Wrapper *WrapperQuery `json:"wrapper,omitempty"`
}

func (s *Query) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s Query) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewQuery() *Query { _ = "STUB: not implemented"; return nil }

type QueryVariant interface {
	QueryCaster() *Query
}

func (s *Query) QueryCaster() *Query { _ = "STUB: not implemented"; return nil }

func (s *Query) IndicesPrivilegesQueryCaster() *IndicesPrivilegesQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *Query) RoleTemplateInlineQueryCaster() *RoleTemplateInlineQuery {
	_ = "STUB: not implemented"
	return nil
}
