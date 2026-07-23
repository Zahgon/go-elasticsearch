package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _sortCombinations struct {
	v types.SortCombinations
}

func NewSortCombinations() *_sortCombinations { _ = "STUB: not implemented"; return nil }

func (u *_sortCombinations) Field(field string) *_sortCombinations {
	_ = "STUB: not implemented"
	return nil
}

func (u *_sortCombinations) SortOptions(sortoptions types.SortOptionsVariant) *_sortCombinations {
	_ = "STUB: not implemented"
	return nil
}

func (u *_sortOptions) SortCombinationsCaster() *types.SortCombinations {
	_ = "STUB: not implemented"
	return nil
}

func (u *_sortCombinations) SortCombinationsCaster() *types.SortCombinations {
	_ = "STUB: not implemented"
	return nil
}
