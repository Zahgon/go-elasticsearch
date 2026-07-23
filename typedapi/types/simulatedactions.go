package types

type SimulatedActions struct {
	Actions []string          `json:"actions"`
	All     *SimulatedActions `json:"all,omitempty"`
	UseAll  bool              `json:"use_all"`
}

func (s *SimulatedActions) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSimulatedActions() *SimulatedActions { _ = "STUB: not implemented"; return nil }

type SimulatedActionsVariant interface {
	SimulatedActionsCaster() *SimulatedActions
}

func (s *SimulatedActions) SimulatedActionsCaster() *SimulatedActions {
	_ = "STUB: not implemented"
	return nil
}
