package types

type Hyperparameter struct {
	AbsoluteImportance *Float64 `json:"absolute_importance,omitempty"`

	Name string `json:"name"`

	RelativeImportance *Float64 `json:"relative_importance,omitempty"`

	Supplied bool `json:"supplied"`

	Value Float64 `json:"value"`
}

func (s *Hyperparameter) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewHyperparameter() *Hyperparameter { _ = "STUB: not implemented"; return nil }
