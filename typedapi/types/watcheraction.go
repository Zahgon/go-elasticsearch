package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/actiontype"
)

type WatcherAction struct {
	ActionType             *actiontype.ActionType `json:"action_type,omitempty"`
	Condition              *WatcherCondition      `json:"condition,omitempty"`
	Email                  *EmailAction           `json:"email,omitempty"`
	Foreach                *string                `json:"foreach,omitempty"`
	Index                  *IndexAction           `json:"index,omitempty"`
	Logging                *LoggingAction         `json:"logging,omitempty"`
	MaxIterations          *int                   `json:"max_iterations,omitempty"`
	Name                   *string                `json:"name,omitempty"`
	Pagerduty              *PagerDutyAction       `json:"pagerduty,omitempty"`
	Slack                  *SlackAction           `json:"slack,omitempty"`
	ThrottlePeriod         Duration               `json:"throttle_period,omitempty"`
	ThrottlePeriodInMillis *int64                 `json:"throttle_period_in_millis,omitempty"`
	Transform              *TransformContainer    `json:"transform,omitempty"`
	Webhook                *WebhookAction         `json:"webhook,omitempty"`
}

func (s *WatcherAction) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewWatcherAction() *WatcherAction { _ = "STUB: not implemented"; return nil }

type WatcherActionVariant interface {
	WatcherActionCaster() *WatcherAction
}

func (s *WatcherAction) WatcherActionCaster() *WatcherAction { _ = "STUB: not implemented"; return nil }
