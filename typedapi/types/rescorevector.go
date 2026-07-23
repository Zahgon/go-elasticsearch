package types

type RescoreVector struct {
	Oversample float32 `json:"oversample"`
}

func (s *RescoreVector) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRescoreVector() *RescoreVector { _ = "STUB: not implemented"; return nil }

type RescoreVectorVariant interface {
	RescoreVectorCaster() *RescoreVector
}

func (s *RescoreVector) RescoreVectorCaster() *RescoreVector { _ = "STUB: not implemented"; return nil }
