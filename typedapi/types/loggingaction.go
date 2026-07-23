package types

type LoggingAction struct {
	Category *string `json:"category,omitempty"`
	Level    *string `json:"level,omitempty"`
	Text     string  `json:"text"`
}

func (s *LoggingAction) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewLoggingAction() *LoggingAction { _ = "STUB: not implemented"; return nil }

type LoggingActionVariant interface {
	LoggingActionCaster() *LoggingAction
}

func (s *LoggingAction) LoggingActionCaster() *LoggingAction { _ = "STUB: not implemented"; return nil }
