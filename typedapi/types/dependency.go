package types

type Dependency struct {
	Field string      `json:"field"`
	Value ScalarValue `json:"value"`
}

func (s *Dependency) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewDependency() *Dependency { _ = "STUB: not implemented"; return nil }

type DependencyVariant interface {
	DependencyCaster() *Dependency
}

func (s *Dependency) DependencyCaster() *Dependency { _ = "STUB: not implemented"; return nil }
