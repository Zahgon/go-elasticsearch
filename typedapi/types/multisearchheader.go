package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/expandwildcard"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/searchtype"
)

type MultisearchHeader struct {
	AllowNoIndices            *bool                           `json:"allow_no_indices,omitempty"`
	AllowPartialSearchResults *bool                           `json:"allow_partial_search_results,omitempty"`
	CcsMinimizeRoundtrips     *bool                           `json:"ccs_minimize_roundtrips,omitempty"`
	ExpandWildcards           []expandwildcard.ExpandWildcard `json:"expand_wildcards,omitempty"`
	IgnoreThrottled           *bool                           `json:"ignore_throttled,omitempty"`

	IgnoreUnavailable *bool    `json:"ignore_unavailable,omitempty"`
	Index             []string `json:"index,omitempty"`
	Preference        *string  `json:"preference,omitempty"`
	ProjectRouting    *string  `json:"project_routing,omitempty"`
	RequestCache      *bool    `json:"request_cache,omitempty"`

	Routing    []string               `json:"routing,omitempty"`
	SearchType *searchtype.SearchType `json:"search_type,omitempty"`

	Slice_ *string `json:"_slice,omitempty"`
}

func (s *MultisearchHeader) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewMultisearchHeader() *MultisearchHeader { _ = "STUB: not implemented"; return nil }

type MultisearchHeaderVariant interface {
	MultisearchHeaderCaster() *MultisearchHeader
}

func (s *MultisearchHeader) MultisearchHeaderCaster() *MultisearchHeader {
	_ = "STUB: not implemented"
	return nil
}

func (s *MultisearchHeader) MsearchRequestItemCaster() *MsearchRequestItem {
	_ = "STUB: not implemented"
	return nil
}

func (s *MultisearchHeader) RequestItemCaster() *RequestItem { _ = "STUB: not implemented"; return nil }
