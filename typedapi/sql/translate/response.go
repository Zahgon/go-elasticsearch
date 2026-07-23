package translate

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Aggregations   map[string]types.Aggregations `json:"aggregations,omitempty"`
	Fields         []types.FieldAndFormat        `json:"fields,omitempty"`
	Query          *types.Query                  `json:"query,omitempty"`
	Size           *int64                        `json:"size,omitempty"`
	Sort           []types.SortCombinations      `json:"sort,omitempty"`
	Source_        types.SourceConfig            `json:"_source,omitempty"`
	TrackTotalHits types.TrackHits               `json:"track_total_hits,omitempty"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }

func (s *Response) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
