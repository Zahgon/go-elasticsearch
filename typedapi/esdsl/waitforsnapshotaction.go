package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _waitForSnapshotAction struct {
	v *types.WaitForSnapshotAction
}

func NewWaitForSnapshotAction(policy string) *_waitForSnapshotAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_waitForSnapshotAction) Policy(policy string) *_waitForSnapshotAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_waitForSnapshotAction) WaitForSnapshotActionCaster() *types.WaitForSnapshotAction {
	_ = "STUB: not implemented"
	return nil
}
