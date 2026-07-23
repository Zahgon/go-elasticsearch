package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _ilmActions struct {
	v *types.IlmActions
}

func NewIlmActions() *_ilmActions { _ = "STUB: not implemented"; return nil }

func (s *_ilmActions) Allocate(allocate types.AllocateActionVariant) *_ilmActions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ilmActions) Delete(delete types.DeleteActionVariant) *_ilmActions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ilmActions) Downsample(downsample types.DownsampleActionVariant) *_ilmActions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ilmActions) Forcemerge(forcemerge types.ForceMergeActionVariant) *_ilmActions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ilmActions) Freeze(freeze types.EmptyObjectVariant) *_ilmActions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ilmActions) Migrate(migrate types.MigrateActionVariant) *_ilmActions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ilmActions) Readonly(readonly types.EmptyObjectVariant) *_ilmActions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ilmActions) Rollover(rollover types.RolloverActionVariant) *_ilmActions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ilmActions) SearchableSnapshot(searchablesnapshot types.SearchableSnapshotActionVariant) *_ilmActions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ilmActions) SetPriority(setpriority types.SetPriorityActionVariant) *_ilmActions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ilmActions) Shrink(shrink types.ShrinkActionVariant) *_ilmActions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ilmActions) Unfollow(unfollow types.EmptyObjectVariant) *_ilmActions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ilmActions) WaitForSnapshot(waitforsnapshot types.WaitForSnapshotActionVariant) *_ilmActions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ilmActions) IlmActionsCaster() *types.IlmActions { _ = "STUB: not implemented"; return nil }
