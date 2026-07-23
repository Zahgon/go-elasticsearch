package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _commandCancelAction struct {
	v *types.CommandCancelAction
}

func NewCommandCancelAction(node string, shard int) *_commandCancelAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_commandCancelAction) AllowPrimary(allowprimary bool) *_commandCancelAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_commandCancelAction) Index(indexname string) *_commandCancelAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_commandCancelAction) Node(node string) *_commandCancelAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_commandCancelAction) Shard(shard int) *_commandCancelAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_commandCancelAction) CommandCancelActionCaster() *types.CommandCancelAction {
	_ = "STUB: not implemented"
	return nil
}
