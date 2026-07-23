package submit

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	Aggregations map[string]types.Aggregations `json:"aggregations,omitempty"`
	Collapse     *types.FieldCollapse          `json:"collapse,omitempty"`

	DocvalueFields []types.FieldAndFormat `json:"docvalue_fields,omitempty"`

	Explain *bool `json:"explain,omitempty"`

	Ext map[string]json.RawMessage `json:"ext,omitempty"`

	Fields []types.FieldAndFormat `json:"fields,omitempty"`

	From      *int             `json:"from,omitempty"`
	Highlight *types.Highlight `json:"highlight,omitempty"`

	IndicesBoost []map[string]types.Float64 `json:"indices_boost,omitempty"`

	Knn []types.KnnSearch `json:"knn,omitempty"`

	MinScore *types.Float64 `json:"min_score,omitempty"`

	Pit        *types.PointInTimeReference `json:"pit,omitempty"`
	PostFilter *types.Query                `json:"post_filter,omitempty"`
	Profile    *bool                       `json:"profile,omitempty"`

	ProjectRouting *string `json:"project_routing,omitempty"`

	Query   *types.Query    `json:"query,omitempty"`
	Rescore []types.Rescore `json:"rescore,omitempty"`

	RuntimeMappings types.RuntimeFields `json:"runtime_mappings,omitempty"`

	ScriptFields map[string]types.ScriptField `json:"script_fields,omitempty"`
	SearchAfter  []types.FieldValue           `json:"search_after,omitempty"`

	SeqNoPrimaryTerm *bool `json:"seq_no_primary_term,omitempty"`

	Size  *int                     `json:"size,omitempty"`
	Slice *types.SlicedScroll      `json:"slice,omitempty"`
	Sort  []types.SortCombinations `json:"sort,omitempty"`

	Source_ types.SourceConfig `json:"_source,omitempty"`

	Stats []string `json:"stats,omitempty"`

	StoredFields []string         `json:"stored_fields,omitempty"`
	Suggest      *types.Suggester `json:"suggest,omitempty"`

	TerminateAfter *int64 `json:"terminate_after,omitempty"`

	Timeout *string `json:"timeout,omitempty"`

	TrackScores *bool `json:"track_scores,omitempty"`

	TrackTotalHits types.TrackHits `json:"track_total_hits,omitempty"`

	Version *bool `json:"version,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (r Request) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
