package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _dependency struct {
	v *types.Dependency
}

func NewDependency(field string) *_dependency { _ = "STUB: not implemented"; return nil }

func (s *_dependency) Field(field string) *_dependency { _ = "STUB: not implemented"; return nil }

func (s *_dependency) Value(scalarvalue types.ScalarValueVariant) *_dependency {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dependency) DependencyCaster() *types.Dependency { _ = "STUB: not implemented"; return nil }
