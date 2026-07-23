package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/actiontype"
)

type _watcherAction struct {
	v *types.WatcherAction
}

func NewWatcherAction() *_watcherAction { _ = "STUB: not implemented"; return nil }

func (s *_watcherAction) ActionType(actiontype actiontype.ActionType) *_watcherAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_watcherAction) Condition(condition types.WatcherConditionVariant) *_watcherAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_watcherAction) Email(email types.EmailActionVariant) *_watcherAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_watcherAction) Foreach(foreach string) *_watcherAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_watcherAction) Index(index types.IndexActionVariant) *_watcherAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_watcherAction) Logging(logging types.LoggingActionVariant) *_watcherAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_watcherAction) MaxIterations(maxiterations int) *_watcherAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_watcherAction) Name(name string) *_watcherAction { _ = "STUB: not implemented"; return nil }

func (s *_watcherAction) Pagerduty(pagerduty types.PagerDutyActionVariant) *_watcherAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_watcherAction) Slack(slack types.SlackActionVariant) *_watcherAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_watcherAction) ThrottlePeriod(duration types.DurationVariant) *_watcherAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_watcherAction) ThrottlePeriodInMillis(durationvalueunitmillis int64) *_watcherAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_watcherAction) Transform(transform types.TransformContainerVariant) *_watcherAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_watcherAction) Webhook(webhook types.WebhookActionVariant) *_watcherAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_watcherAction) WatcherActionCaster() *types.WatcherAction {
	_ = "STUB: not implemented"
	return nil
}
