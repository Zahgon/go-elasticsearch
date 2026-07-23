package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _selectOption struct {
	v *types.SelectOption
}

func NewSelectOption(label string) *_selectOption { _ = "STUB: not implemented"; return nil }

func (s *_selectOption) Label(label string) *_selectOption { _ = "STUB: not implemented"; return nil }

func (s *_selectOption) Value(scalarvalue types.ScalarValueVariant) *_selectOption {
	_ = "STUB: not implemented"
	return nil
}

func (s *_selectOption) SelectOptionCaster() *types.SelectOption {
	_ = "STUB: not implemented"
	return nil
}
