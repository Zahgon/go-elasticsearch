package types

import (
	"encoding/json"
)

type RoleQueryContainer struct {
	AdditionalRoleQueryContainerProperty map[string]json.RawMessage `json:"-"`

	Bool *BoolQuery `json:"bool,omitempty"`

	Exists *ExistsQuery `json:"exists,omitempty"`

	Ids *IdsQuery `json:"ids,omitempty"`

	Match map[string]MatchQuery `json:"match,omitempty"`

	MatchAll *MatchAllQuery `json:"match_all,omitempty"`

	Prefix map[string]PrefixQuery `json:"prefix,omitempty"`

	Range map[string]RangeQuery `json:"range,omitempty"`

	SimpleQueryString *SimpleQueryStringQuery `json:"simple_query_string,omitempty"`

	Term map[string]TermQuery `json:"term,omitempty"`

	Terms *TermsQuery `json:"terms,omitempty"`

	Wildcard map[string]WildcardQuery `json:"wildcard,omitempty"`
}

func (s *RoleQueryContainer) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s RoleQueryContainer) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewRoleQueryContainer() *RoleQueryContainer { _ = "STUB: not implemented"; return nil }

type RoleQueryContainerVariant interface {
	RoleQueryContainerCaster() *RoleQueryContainer
}

func (s *RoleQueryContainer) RoleQueryContainerCaster() *RoleQueryContainer {
	_ = "STUB: not implemented"
	return nil
}
