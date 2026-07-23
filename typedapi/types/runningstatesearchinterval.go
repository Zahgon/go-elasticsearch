package types

type RunningStateSearchInterval struct {
	End Duration `json:"end,omitempty"`

	EndMs int64 `json:"end_ms"`

	Start Duration `json:"start,omitempty"`

	StartMs int64 `json:"start_ms"`
}

func (s *RunningStateSearchInterval) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewRunningStateSearchInterval() *RunningStateSearchInterval {
	_ = "STUB: not implemented"
	return nil
}
