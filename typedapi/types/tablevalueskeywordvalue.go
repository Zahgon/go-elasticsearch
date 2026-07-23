package types

type TableValuesKeywordValue []string

type TableValuesKeywordValueVariant interface {
	TableValuesKeywordValueCaster() *TableValuesKeywordValue
}
