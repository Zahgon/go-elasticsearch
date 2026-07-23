package query

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	Columnar *bool `json:"columnar,omitempty"`

	Filter *types.Query `json:"filter,omitempty"`

	IncludeCcsMetadata *bool `json:"include_ccs_metadata,omitempty"`

	IncludeExecutionMetadata *bool `json:"include_execution_metadata,omitempty"`

	Locale *string `json:"locale,omitempty"`

	Params types.ESQLParams `json:"params,omitempty"`

	Profile *bool `json:"profile,omitempty"`

	ProjectRouting *string `json:"project_routing,omitempty"`

	Query string `json:"query"`

	Tables map[string]map[string]types.TableValuesContainer `json:"tables,omitempty"`

	TimeZone *string `json:"time_zone,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
