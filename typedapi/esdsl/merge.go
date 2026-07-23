package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _merge struct {
	v *types.Merge
}

func NewMerge() *_merge { _ = "STUB: not implemented"; return nil }

func (s *_merge) Scheduler(scheduler types.MergeSchedulerVariant) *_merge {
	_ = "STUB: not implemented"
	return nil
}

func (s *_merge) MergeCaster() *types.Merge { _ = "STUB: not implemented"; return nil }
