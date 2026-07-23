package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/filteringpolicy"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/filteringrulerule"
)

type FilteringRule struct {
	CreatedAt DateTime                            `json:"created_at,omitempty"`
	Field     string                              `json:"field"`
	Id        string                              `json:"id"`
	Order     int                                 `json:"order"`
	Policy    filteringpolicy.FilteringPolicy     `json:"policy"`
	Rule      filteringrulerule.FilteringRuleRule `json:"rule"`
	UpdatedAt DateTime                            `json:"updated_at,omitempty"`
	Value     string                              `json:"value"`
}

func (s *FilteringRule) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewFilteringRule() *FilteringRule { _ = "STUB: not implemented"; return nil }

type FilteringRuleVariant interface {
	FilteringRuleCaster() *FilteringRule
}

func (s *FilteringRule) FilteringRuleCaster() *FilteringRule { _ = "STUB: not implemented"; return nil }
