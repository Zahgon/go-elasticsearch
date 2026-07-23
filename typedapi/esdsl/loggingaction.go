package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _loggingAction struct {
	v *types.LoggingAction
}

func NewLoggingAction(text string) *_loggingAction { _ = "STUB: not implemented"; return nil }

func (s *_loggingAction) Category(category string) *_loggingAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_loggingAction) Level(level string) *_loggingAction { _ = "STUB: not implemented"; return nil }

func (s *_loggingAction) Text(text string) *_loggingAction { _ = "STUB: not implemented"; return nil }

func (s *_loggingAction) LoggingActionCaster() *types.LoggingAction {
	_ = "STUB: not implemented"
	return nil
}
