package types

type FileSettingsIndicatorDetails struct {
	FailureStreak     int64  `json:"failure_streak"`
	MostRecentFailure string `json:"most_recent_failure"`
}

func (s *FileSettingsIndicatorDetails) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewFileSettingsIndicatorDetails() *FileSettingsIndicatorDetails {
	_ = "STUB: not implemented"
	return nil
}
