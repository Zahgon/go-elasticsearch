package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _watch struct {
	v *types.Watch
}

func NewWatch(condition types.WatcherConditionVariant, input types.WatcherInputVariant, trigger types.TriggerContainerVariant) *_watch {
	_ = "STUB: not implemented"
	return nil
}

func (s *_watch) Actions(actions map[string]types.WatcherAction) *_watch {
	_ = "STUB: not implemented"
	return nil
}

func (s *_watch) AddAction(key string, value types.WatcherActionVariant) *_watch {
	_ = "STUB: not implemented"
	return nil
}

func (s *_watch) Condition(condition types.WatcherConditionVariant) *_watch {
	_ = "STUB: not implemented"
	return nil
}

func (s *_watch) Input(input types.WatcherInputVariant) *_watch {
	_ = "STUB: not implemented"
	return nil
}

func (s *_watch) Metadata(metadata types.MetadataVariant) *_watch {
	_ = "STUB: not implemented"
	return nil
}

func (s *_watch) Status(status types.WatchStatusVariant) *_watch {
	_ = "STUB: not implemented"
	return nil
}

func (s *_watch) ThrottlePeriod(duration types.DurationVariant) *_watch {
	_ = "STUB: not implemented"
	return nil
}

func (s *_watch) ThrottlePeriodInMillis(durationvalueunitmillis int64) *_watch {
	_ = "STUB: not implemented"
	return nil
}

func (s *_watch) Transform(transform types.TransformContainerVariant) *_watch {
	_ = "STUB: not implemented"
	return nil
}

func (s *_watch) Trigger(trigger types.TriggerContainerVariant) *_watch {
	_ = "STUB: not implemented"
	return nil
}

func (s *_watch) WatchCaster() *types.Watch { _ = "STUB: not implemented"; return nil }
