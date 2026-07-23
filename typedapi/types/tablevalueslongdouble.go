package types

type TableValuesLongDouble []Float64

type TableValuesLongDoubleVariant interface {
	TableValuesLongDoubleCaster() *TableValuesLongDouble
}
