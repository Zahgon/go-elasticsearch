package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _slackAction struct {
	v *types.SlackAction
}

func NewSlackAction(message types.SlackMessageVariant) *_slackAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_slackAction) Account(account string) *_slackAction { _ = "STUB: not implemented"; return nil }

func (s *_slackAction) Message(message types.SlackMessageVariant) *_slackAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_slackAction) SlackActionCaster() *types.SlackAction {
	_ = "STUB: not implemented"
	return nil
}
