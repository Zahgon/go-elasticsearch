package types

type TableValuesLongValue []int64

type TableValuesLongValueVariant interface {
	TableValuesLongValueCaster() *TableValuesLongValue
}
