package types

type IlmActions struct {
	Allocate *AllocateAction `json:"allocate,omitempty"`

	Delete *DeleteAction `json:"delete,omitempty"`

	Downsample *DownsampleAction `json:"downsample,omitempty"`

	Forcemerge *ForceMergeAction `json:"forcemerge,omitempty"`

	Freeze *EmptyObject `json:"freeze,omitempty"`

	Migrate *MigrateAction `json:"migrate,omitempty"`

	Readonly *EmptyObject `json:"readonly,omitempty"`

	Rollover *RolloverAction `json:"rollover,omitempty"`

	SearchableSnapshot *SearchableSnapshotAction `json:"searchable_snapshot,omitempty"`

	SetPriority *SetPriorityAction `json:"set_priority,omitempty"`

	Shrink *ShrinkAction `json:"shrink,omitempty"`

	Unfollow *EmptyObject `json:"unfollow,omitempty"`

	WaitForSnapshot *WaitForSnapshotAction `json:"wait_for_snapshot,omitempty"`
}

func NewIlmActions() *IlmActions { _ = "STUB: not implemented"; return nil }

type IlmActionsVariant interface {
	IlmActionsCaster() *IlmActions
}

func (s *IlmActions) IlmActionsCaster() *IlmActions { _ = "STUB: not implemented"; return nil }
