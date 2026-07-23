package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _searchInput struct {
	v *types.SearchInput
}

func NewSearchInput(request types.SearchInputRequestDefinitionVariant) *_searchInput {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchInput) Extract(extracts ...string) *_searchInput {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchInput) Request(request types.SearchInputRequestDefinitionVariant) *_searchInput {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchInput) Timeout(duration types.DurationVariant) *_searchInput {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchInput) WatcherInputCaster() *types.WatcherInput {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchInput) SearchInputCaster() *types.SearchInput {
	_ = "STUB: not implemented"
	return nil
}
