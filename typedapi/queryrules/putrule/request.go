package putrule

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/queryruletype"
)

type Request struct {
	Actions types.QueryRuleActions `json:"actions"`

	Criteria []types.QueryRuleCriteria `json:"criteria"`
	Priority *int                      `json:"priority,omitempty"`

	Type queryruletype.QueryRuleType `json:"type"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
