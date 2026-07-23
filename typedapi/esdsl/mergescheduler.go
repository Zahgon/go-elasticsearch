package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _mergeScheduler struct {
	v *types.MergeScheduler
}

func NewMergeScheduler() *_mergeScheduler { _ = "STUB: not implemented"; return nil }

func (s *_mergeScheduler) MaxMergeCount(stringifiedinteger types.StringifiedintegerVariant) *_mergeScheduler {
	_ = "STUB: not implemented"
	return nil
}

func (s *_mergeScheduler) MaxThreadCount(stringifiedinteger types.StringifiedintegerVariant) *_mergeScheduler {
	_ = "STUB: not implemented"
	return nil
}

func (s *_mergeScheduler) MergeSchedulerCaster() *types.MergeScheduler {
	_ = "STUB: not implemented"
	return nil
}
