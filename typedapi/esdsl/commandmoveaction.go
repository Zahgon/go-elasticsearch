package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _commandMoveAction struct {
	v *types.CommandMoveAction
}

func NewCommandMoveAction(fromnode string, shard int, tonode string) *_commandMoveAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_commandMoveAction) FromNode(fromnode string) *_commandMoveAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_commandMoveAction) Index(indexname string) *_commandMoveAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_commandMoveAction) Shard(shard int) *_commandMoveAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_commandMoveAction) ToNode(tonode string) *_commandMoveAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_commandMoveAction) CommandMoveActionCaster() *types.CommandMoveAction {
	_ = "STUB: not implemented"
	return nil
}
