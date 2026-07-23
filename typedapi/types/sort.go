package types

type Sort []SortCombinations

type SortVariant interface {
	SortCaster() *Sort
}
