package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/expandwildcard"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/searchtype"
)

type _multisearchHeader struct {
	v *types.MultisearchHeader
}

func NewMultisearchHeader() *_multisearchHeader { _ = "STUB: not implemented"; return nil }

func (s *_multisearchHeader) AllowNoIndices(allownoindices bool) *_multisearchHeader {
	_ = "STUB: not implemented"
	return nil
}

func (s *_multisearchHeader) AllowPartialSearchResults(allowpartialsearchresults bool) *_multisearchHeader {
	_ = "STUB: not implemented"
	return nil
}

func (s *_multisearchHeader) CcsMinimizeRoundtrips(ccsminimizeroundtrips bool) *_multisearchHeader {
	_ = "STUB: not implemented"
	return nil
}

func (s *_multisearchHeader) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *_multisearchHeader {
	_ = "STUB: not implemented"
	return nil
}

func (s *_multisearchHeader) IgnoreThrottled(ignorethrottled bool) *_multisearchHeader {
	_ = "STUB: not implemented"
	return nil
}

func (s *_multisearchHeader) IgnoreUnavailable(ignoreunavailable bool) *_multisearchHeader {
	_ = "STUB: not implemented"
	return nil
}

func (s *_multisearchHeader) Index(indices ...string) *_multisearchHeader {
	_ = "STUB: not implemented"
	return nil
}

func (s *_multisearchHeader) Preference(preference string) *_multisearchHeader {
	_ = "STUB: not implemented"
	return nil
}

func (s *_multisearchHeader) ProjectRouting(projectrouting string) *_multisearchHeader {
	_ = "STUB: not implemented"
	return nil
}

func (s *_multisearchHeader) RequestCache(requestcache bool) *_multisearchHeader {
	_ = "STUB: not implemented"
	return nil
}

func (s *_multisearchHeader) Routing(routings ...string) *_multisearchHeader {
	_ = "STUB: not implemented"
	return nil
}

func (s *_multisearchHeader) SearchType(searchtype searchtype.SearchType) *_multisearchHeader {
	_ = "STUB: not implemented"
	return nil
}

func (s *_multisearchHeader) Slice_(slice_ string) *_multisearchHeader {
	_ = "STUB: not implemented"
	return nil
}

func (s *_multisearchHeader) MultisearchHeaderCaster() *types.MultisearchHeader {
	_ = "STUB: not implemented"
	return nil
}
