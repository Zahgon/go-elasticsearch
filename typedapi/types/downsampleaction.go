package types

type DownsampleAction struct {
	FixedInterval string   `json:"fixed_interval"`
	WaitTimeout   Duration `json:"wait_timeout,omitempty"`
}

func (s *DownsampleAction) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewDownsampleAction() *DownsampleAction { _ = "STUB: not implemented"; return nil }

type DownsampleActionVariant interface {
	DownsampleActionCaster() *DownsampleAction
}

func (s *DownsampleAction) DownsampleActionCaster() *DownsampleAction {
	_ = "STUB: not implemented"
	return nil
}
