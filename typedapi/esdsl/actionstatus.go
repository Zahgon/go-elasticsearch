package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _actionStatus struct {
	v *types.ActionStatus
}

func NewActionStatus(ack types.AcknowledgeStateVariant) *_actionStatus {
	_ = "STUB: not implemented"
	return nil
}

func (s *_actionStatus) Ack(ack types.AcknowledgeStateVariant) *_actionStatus {
	_ = "STUB: not implemented"
	return nil
}

func (s *_actionStatus) LastExecution(lastexecution types.ExecutionStateVariant) *_actionStatus {
	_ = "STUB: not implemented"
	return nil
}

func (s *_actionStatus) LastSuccessfulExecution(lastsuccessfulexecution types.ExecutionStateVariant) *_actionStatus {
	_ = "STUB: not implemented"
	return nil
}

func (s *_actionStatus) LastThrottle(lastthrottle types.ThrottleStateVariant) *_actionStatus {
	_ = "STUB: not implemented"
	return nil
}

func (s *_actionStatus) ActionStatusCaster() *types.ActionStatus {
	_ = "STUB: not implemented"
	return nil
}
