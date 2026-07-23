package types

import (
	"encoding/json"
)

type SearchRequestBody struct {
	Aggregations map[string]Aggregations `json:"aggregations,omitempty"`

	Collapse *FieldCollapse `json:"collapse,omitempty"`

	DocvalueFields []FieldAndFormat `json:"docvalue_fields,omitempty"`

	Explain *bool `json:"explain,omitempty"`

	Ext map[string]json.RawMessage `json:"ext,omitempty"`

	Fields []FieldAndFormat `json:"fields,omitempty"`

	From *int `json:"from,omitempty"`

	Highlight *Highlight `json:"highlight,omitempty"`

	IndicesBoost []map[string]Float64 `json:"indices_boost,omitempty"`

	Knn []KnnSearch `json:"knn,omitempty"`

	MinScore *Float64 `json:"min_score,omitempty"`

	Pit *PointInTimeReference `json:"pit,omitempty"`

	PostFilter *Query `json:"post_filter,omitempty"`

	Profile *bool `json:"profile,omitempty"`

	Query *Query `json:"query,omitempty"`

	Rank *RankContainer `json:"rank,omitempty"`

	Rescore []Rescore `json:"rescore,omitempty"`

	Retriever *RetrieverContainer `json:"retriever,omitempty"`

	RuntimeMappings RuntimeFields `json:"runtime_mappings,omitempty"`

	ScriptFields map[string]ScriptField `json:"script_fields,omitempty"`

	SearchAfter []FieldValue `json:"search_after,omitempty"`

	SeqNoPrimaryTerm *bool `json:"seq_no_primary_term,omitempty"`

	Size *int `json:"size,omitempty"`

	Slice *SlicedScroll `json:"slice,omitempty"`

	Sort []SortCombinations `json:"sort,omitempty"`

	Source_ SourceConfig `json:"_source,omitempty"`

	Stats []string `json:"stats,omitempty"`

	StoredFields []string `json:"stored_fields,omitempty"`

	Suggest *Suggester `json:"suggest,omitempty"`

	TerminateAfter *int64 `json:"terminate_after,omitempty"`

	Timeout *string `json:"timeout,omitempty"`

	TrackScores *bool `json:"track_scores,omitempty"`

	TrackTotalHits TrackHits `json:"track_total_hits,omitempty"`

	Version *bool `json:"version,omitempty"`
}

func (s *SearchRequestBody) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSearchRequestBody() *SearchRequestBody { _ = "STUB: not implemented"; return nil }

type SearchRequestBodyVariant interface {
	SearchRequestBodyCaster() *SearchRequestBody
}

func (s *SearchRequestBody) SearchRequestBodyCaster() *SearchRequestBody {
	_ = "STUB: not implemented"
	return nil
}

func (s *SearchRequestBody) MsearchRequestItemCaster() *MsearchRequestItem {
	_ = "STUB: not implemented"
	return nil
}

func (s *SearchRequestBody) ScriptSourceCaster() *ScriptSource {
	_ = "STUB: not implemented"
	return nil
}
