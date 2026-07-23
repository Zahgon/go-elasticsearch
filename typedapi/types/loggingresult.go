package types

type LoggingResult struct {
	LoggedText string `json:"logged_text"`
}

func (s *LoggingResult) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewLoggingResult() *LoggingResult { _ = "STUB: not implemented"; return nil }
