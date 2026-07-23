package types

type DownsamplingRound struct {
	After Duration `json:"after"`

	FixedInterval string `json:"fixed_interval"`
}

func (s *DownsamplingRound) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewDownsamplingRound() *DownsamplingRound { _ = "STUB: not implemented"; return nil }

type DownsamplingRoundVariant interface {
	DownsamplingRoundCaster() *DownsamplingRound
}

func (s *DownsamplingRound) DownsamplingRoundCaster() *DownsamplingRound {
	_ = "STUB: not implemented"
	return nil
}
