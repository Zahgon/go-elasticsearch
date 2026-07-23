package searchmvt

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gridaggregationtype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gridtype"
)

type Request struct {
	Aggs map[string]types.Aggregations `json:"aggs,omitempty"`

	Buffer *int `json:"buffer,omitempty"`

	ExactBounds *bool `json:"exact_bounds,omitempty"`

	Extent *int `json:"extent,omitempty"`

	Fields []string `json:"fields,omitempty"`

	GridAgg *gridaggregationtype.GridAggregationType `json:"grid_agg,omitempty"`

	GridPrecision *int `json:"grid_precision,omitempty"`

	GridType *gridtype.GridType `json:"grid_type,omitempty"`

	ProjectRouting *string `json:"project_routing,omitempty"`

	Query *types.Query `json:"query,omitempty"`

	RuntimeMappings types.RuntimeFields `json:"runtime_mappings,omitempty"`

	Size *int `json:"size,omitempty"`

	Sort []types.SortCombinations `json:"sort,omitempty"`

	TrackTotalHits types.TrackHits `json:"track_total_hits,omitempty"`

	WithLabels *bool `json:"with_labels,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
