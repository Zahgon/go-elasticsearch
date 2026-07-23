package types

type TableValuesContainer struct {
	Float64 [][]Float64 `json:"double,omitempty"`
	Int     [][]int     `json:"integer,omitempty"`
	Int64   [][]int64   `json:"long,omitempty"`
	Keyword [][]string  `json:"keyword,omitempty"`
}

func NewTableValuesContainer() *TableValuesContainer { _ = "STUB: not implemented"; return nil }

type TableValuesContainerVariant interface {
	TableValuesContainerCaster() *TableValuesContainer
}

func (s *TableValuesContainer) TableValuesContainerCaster() *TableValuesContainer {
	_ = "STUB: not implemented"
	return nil
}
