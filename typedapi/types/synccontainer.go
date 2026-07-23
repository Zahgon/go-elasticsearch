package types

type SyncContainer struct {
	Time *TimeSync `json:"time,omitempty"`
}

func NewSyncContainer() *SyncContainer { _ = "STUB: not implemented"; return nil }

type SyncContainerVariant interface {
	SyncContainerCaster() *SyncContainer
}

func (s *SyncContainer) SyncContainerCaster() *SyncContainer { _ = "STUB: not implemented"; return nil }
