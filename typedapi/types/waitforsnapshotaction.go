package types

type WaitForSnapshotAction struct {
	Policy string `json:"policy"`
}

func (s *WaitForSnapshotAction) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewWaitForSnapshotAction() *WaitForSnapshotAction { _ = "STUB: not implemented"; return nil }

type WaitForSnapshotActionVariant interface {
	WaitForSnapshotActionCaster() *WaitForSnapshotAction
}

func (s *WaitForSnapshotAction) WaitForSnapshotActionCaster() *WaitForSnapshotAction {
	_ = "STUB: not implemented"
	return nil
}
