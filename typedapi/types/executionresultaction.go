package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/actionstatusoptions"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/actiontype"
)

type ExecutionResultAction struct {
	Email     *EmailResult                            `json:"email,omitempty"`
	Error     *ErrorCause                             `json:"error,omitempty"`
	Id        string                                  `json:"id"`
	Index     *IndexResult                            `json:"index,omitempty"`
	Logging   *LoggingResult                          `json:"logging,omitempty"`
	Pagerduty *PagerDutyResult                        `json:"pagerduty,omitempty"`
	Reason    *string                                 `json:"reason,omitempty"`
	Slack     *SlackResult                            `json:"slack,omitempty"`
	Status    actionstatusoptions.ActionStatusOptions `json:"status"`
	Type      actiontype.ActionType                   `json:"type"`
	Webhook   *WebhookResult                          `json:"webhook,omitempty"`
}

func (s *ExecutionResultAction) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewExecutionResultAction() *ExecutionResultAction { _ = "STUB: not implemented"; return nil }
