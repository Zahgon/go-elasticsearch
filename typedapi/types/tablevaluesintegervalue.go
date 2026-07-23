package types

type TableValuesIntegerValue []int

type TableValuesIntegerValueVariant interface {
	TableValuesIntegerValueCaster() *TableValuesIntegerValue
}
