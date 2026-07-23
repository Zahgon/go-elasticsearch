package types

type SortCombinations any

type SortCombinationsVariant interface {
	SortCombinationsCaster() *SortCombinations
}
