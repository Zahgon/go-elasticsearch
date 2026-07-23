package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _sortOptions struct {
	v *types.SortOptions
}

func NewSortOptions() *_sortOptions { _ = "STUB: not implemented"; return nil }

func (s *_sortOptions) Doc_(doc_ types.ScoreSortVariant) *_sortOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sortOptions) GeoDistance_(geodistance_ types.GeoDistanceSortVariant) *_sortOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sortOptions) Score_(score_ types.ScoreSortVariant) *_sortOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sortOptions) Script_(script_ types.ScriptSortVariant) *_sortOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sortOptions) SortOptions(sortoptions map[string]types.FieldSort) *_sortOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sortOptions) AddSortOption(key string, value types.FieldSortVariant) *_sortOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sortOptions) SortOptionsCaster() *types.SortOptions {
	_ = "STUB: not implemented"
	return nil
}
