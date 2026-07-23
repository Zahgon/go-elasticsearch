package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _tableValuesContainer struct {
	v *types.TableValuesContainer
}

func NewTableValuesContainer() *_tableValuesContainer { _ = "STUB: not implemented"; return nil }

func (s *_tableValuesContainer) Float64(float64s ...[]types.Float64) *_tableValuesContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_tableValuesContainer) Int(ints ...[]int) *_tableValuesContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_tableValuesContainer) Int64(int64s ...[]int64) *_tableValuesContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_tableValuesContainer) Keyword(keywords ...[]string) *_tableValuesContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_tableValuesContainer) TableValuesContainerCaster() *types.TableValuesContainer {
	_ = "STUB: not implemented"
	return nil
}
