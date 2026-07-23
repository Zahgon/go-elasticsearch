package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _simulatedActions struct {
	v *types.SimulatedActions
}

func NewSimulatedActions(all types.SimulatedActionsVariant, useall bool) *_simulatedActions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_simulatedActions) Actions(actions ...string) *_simulatedActions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_simulatedActions) All(all types.SimulatedActionsVariant) *_simulatedActions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_simulatedActions) UseAll(useall bool) *_simulatedActions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_simulatedActions) SimulatedActionsCaster() *types.SimulatedActions {
	_ = "STUB: not implemented"
	return nil
}
